#!/usr/bin/env just --justfile
# list of available commands
default:
  @just --list

regFile := "./registries"
tagCommit := `[ -d ".git" ] && (git rev-list --abbrev-commit --tags --max-count=1) || true`
lastVersion := `[ -d ".git" ] && (git describe --tags --abbrev=0 2>/dev/null || git symbolic-ref -q --short HEAD) || true`
lastBranchName := `[ -d ".git" ] && (git describe --tags --exact-match HEAD 2>/dev/null || git symbolic-ref -q --short HEAD) || true`
goVersion := `if [ -f "go.mod" ]; then grep -oP '^go\s+\K\d+(\.\d+)?' go.mod; else go version | sed -n 's/.*go\([0-9.]*\).*/\1/p'; fi`
serviceName := `[ -d ".git" ] && (git remote get-url origin 2>/dev/null | sed -E 's#.*/([^/]+)\.git#\1#; s#.*/([^/]+)#\1#') || true`
repoRemoteURL := `[ -d ".git" ] && (git config --get remote.origin.url) || true`
githubOrigin := `[ -d ".git" ] && (git config --get remote.origin.url | grep github.com || echo "") || true`

# prints package informations
info:
    @printf "\033[1mPackage information:\033[0m\n"
    @printf "%-12s\t%s\n" "url:" "{{repoRemoteURL}}"
    @printf "%-12s\t%s\n" "service:" "{{serviceName}}"
    @printf "%-12s\t%s\n" "go-version:" "{{goVersion}}"
    @printf "%-12s\t%s\n" "current-tag:" "{{lastBranchName}}"
    @printf "\n\033[1mRegistries:\033[0m\n"
    @just registries

# clean build directory
clean:
    @echo "clean bin directory..."
    @[ -d "./bin" ] && rm -r ./bin && echo "bin directory cleaned" || true

# clean and build project
build: clean
    go build -o ./bin/service -ldflags="-s -w" -ldflags="-X 'main.Version={{lastBranchName}}' -X 'main.BuildDate=$(date -u '+%Y-%m-%d %H:%M:%S')'" ./cmd

# build and compress binary
upx: build
    upx --best --lzma bin/service

# run go tests
test:
    #!/usr/bin/env bash
    if which gotestsum &> /dev/null; then
        gotestsum --hide-summary="skipped" --junitfile "testout/results.xml" --format "standard-verbose" -- -coverprofile="testout/coverage.out" -race -parallel 1 ./...
        go tool cover -func="testout/coverage.out"
    else
        echo "gotestsum not found!"
        echo "check: https://github.com/gotestyourself/gotestsum#install"
    fi

# run golang linter and fix it
linter fix="false":
    #!/usr/bin/env bash
    if which golangci-lint &> /dev/null; then
        if [[ "{{ fix }}" == "true" ]];then
            golangci-lint run --deadline=20m --concurrency 1 --fix
        else
            golangci-lint run --deadline=20m --concurrency 1
        fi
    else
        echo "linter not found!"
        echo "check: https://golangci-lint.run/usage/install"
    fi

# build service docker image
buildimg tag="auto" cache="true" latest="true":
    #!/usr/bin/env bash
    echo "build service docker image..."
    just info
    imgtag="{{tag}}"
    if [[ "$imgtag" == "auto" ]];then
        imgtag="{{lastBranchName}}"
    fi
    echo "build service docker image..."
    if [ ! -f "{{regFile}}" ]; then
        echo "Error: registries file {{regFile}} not found."
        exit 1
    fi
    mapfile -t registries < "{{regFile}}"
    for registry in "${registries[@]}"; do
        declare -a regparts=($(echo "$registry" | tr ';' ' '))
        reg_host="${regparts[0]}"
        reg_prefix="${regparts[1]}"
        echo "building image: $reg_host/$reg_prefix/{{serviceName}}:$imgtag"
        if [[ "{{cache}}" == "true" ]];then
            docker buildx build -t "$reg_host/$reg_prefix/{{serviceName}}:$imgtag" -f Dockerfile --build-arg GITHUB_TOKEN="$GITHUB_TOKEN" --build-arg GO_VERSION="{{goVersion}}" --build-arg GITHUB_ORIGIN="{{githubOrigin}}" .
        else
            docker buildx build -t "$reg_host/$reg_prefix/{{serviceName}}:$imgtag" -f Dockerfile --build-arg GITHUB_TOKEN="$GITHUB_TOKEN" --build-arg GO_VERSION="{{goVersion}}" --build-arg GITHUB_ORIGIN="{{githubOrigin}}" --no-cache .
        fi
        if [[ "{{ latest }}" == "true" ]];then
            echo "building image: $reg_host/$reg_prefix/{{serviceName}}:latest"
            docker tag "$reg_host/$reg_prefix/{{serviceName}}:$imgtag" "$reg_host/$reg_prefix/{{serviceName}}:latest"
        fi
    done


# build service docker image from built binary
buildimg-bin tag="auto" latest="true": upx
    #!/usr/bin/env bash
    echo "build service docker image from binary..."
    just info
    imgtag="{{tag}}"
    if [[ "$imgtag" == "auto" ]];then
        imgtag="{{lastBranchName}}"
    fi
    if [ ! -f "{{regFile}}" ]; then
        echo "Error: registries file {{regFile}} not found."
        exit 1
    fi
    mapfile -t registries < "{{regFile}}"
    for registry in "${registries[@]}"; do
        declare -a regparts=($(echo "$registry" | tr ';' ' '))
        reg_host="${regparts[0]}"
        reg_prefix="${regparts[1]}"
        echo "building $reg_host/$reg_prefix/{{serviceName}}:$imgtag"
        docker buildx build -t "$reg_host/$reg_prefix/{{serviceName}}:$imgtag" -f binservice.Dockerfile .
        if [[ "{{ latest }}" == "true" ]];then
            echo "building image: $reg_host/$reg_prefix/{{serviceName}}:latest"
            docker tag "$reg_host/$reg_prefix/{{serviceName}}:$imgtag" "$reg_host/$reg_prefix/{{serviceName}}:latest"
        fi
    done

# publish builds and pushes docker image to registry
publishimg tag="auto" latest="true":
    just buildimg "{{tag}}" true "{{latest}}"
    just pushimg "{{tag}}" "{{latest}}"


# push docker image to registry with tag
pushimg tag="auto" latest="true":
    #!/usr/bin/env bash
    echo "push docker images..."
    imgtag="{{tag}}"
    if [[ "$imgtag" == "auto" ]];then
        imgtag="{{lastBranchName}}"
    fi
    if [ ! -f "{{regFile}}" ]; then
        echo "Error: registries file {{regFile}} not found."
        exit 1
    fi
    mapfile -t registries < "{{regFile}}"
    for registry in "${registries[@]}"; do
        declare -a regparts=($(echo "$registry" | tr ';' ' '))
        reg_host="${regparts[0]}"
        reg_prefix="${regparts[1]}"
        username_env="${regparts[2]}"
        password_env="${regparts[3]}"

        reg_username="${!username_env}"
        reg_password="${!password_env}"
        echo "push image: $reg_host/$reg_prefix/{{serviceName}}:$imgtag"
        docker login "$reg_host" -u "$reg_username" -p "$reg_password"
        docker push "$reg_host/$reg_prefix/{{serviceName}}:$imgtag"
        if [[ "{{ latest }}" == "true" ]];then
            docker push "$reg_host/$reg_prefix/{{serviceName}}:latest"
        fi
    done

# release using goreleaser
gorelease local="false":
    #!/usr/bin/env bash
    echo "run go releaser..."
    if which goreleaser&> /dev/null; then
        if [[ "{{ local }}" == "true" ]];then
            goreleaser release --snapshot --clean
        else
            goreleaser release --clean
        fi
    else
        echo "goreleaser not found!"
        echo "check: https://goreleaser.com/install"
    fi

# generate changelog file
gen-changelog:
    #!/usr/bin/env bash
    if which git-chglog &> /dev/null; then
        if [ -d ".chglog" ]; then
            git-chglog -o "CHANGELOG.md"
        else
            git-chglog --init
            just gen-changelog
        fi
    else
        echo "git-changelog not found!"
        echo "check: https://github.com/git-chglog/git-chglog#installation"
    fi

# release target=major/minor/patch
release target="patch" gorelease="false" publishimg="false":
    #!/usr/bin/env bash
    [[ "{{target}}" =~ "^(major|minor|patch)$" ]] || (echo "invalid target: {{target}}" && echo "target should be major/minor/patch")
    is_version() {
         local pattern="^v[0-9]+\.[0-9]+\.[0-9]+$"
         [[ $1 =~ $pattern ]]
    }
    new_version="v1.0.0"
    if is_version "{{lastVersion}}"; then
        major=$(echo "{{lastVersion}}" | cut -d '.' -f 1 | sed 's/v//')
        minor=$(echo "{{lastVersion}}" | cut -d '.' -f 2)
        patch=$(echo "{{lastVersion}}" | cut -d '.' -f 3)
        case "{{target}}" in
            "major")
                new_major=$((major + 1))
                new_version="v$new_major.0.0"
                ;;
            "minor")
                new_minor=$((minor + 1))
                new_version="v$major.$new_minor.0"
                ;;
            "patch")
                new_patch=$((patch + 1))
                new_version="v$major.$minor.$new_patch"
                ;;
            *)
                ;;
        esac
    fi
    echo "release version: $new_version"
    if which git-chglog &> /dev/null; then
        git-chglog -o CHANGELOG.md --next-tag "$new_version"
    else
        echo "git-changelog not found!"
        echo "check: https://github.com/git-chglog/git-chglog#installation"
        exit 1
    fi
    git add -A && git commit -m "release $new_version"
    git tag -a "$new_version" -m "release $new_version"
    git push --follow-tags

    if [[ "{{ gorelease }}" == "true" ]];then
        just gorelease
    fi
    if [[ "{{ publishimg }}" == "true" ]];then
        just publishimg auto false
    fi

# removes last tag from local and remote
rm-last-tag:
    #!/usr/bin/env bash
    is_version() {
      local pattern="^v[0-9]+\.[0-9]+\.[0-9]+$"
      [[ $1 =~ $pattern ]]
    }
    if is_version "{{lastVersion}}"; then
        echo "remove latest tag on local..."
        git tag -d "{{lastVersion}}"
        echo "remove latest tag on remote..."
        git push --delete origin "{{lastVersion}}" || echo "tag not exists on remote"
    fi


# prints list of all tags
list-tags:
    git tag -l --sort=v:refname

# prints list of all registries
registries:
    #!/usr/bin/env bash
    if [ -f "{{regFile}}" ]; then
        mapfile -t registries < "{{regFile}}"
        for registry in "${registries[@]}"; do
            declare -a regparts=($(echo "$registry" | tr ';' ' '))
            reg_host="${regparts[0]}"
            reg_prefix="${regparts[1]}"
            echo "$reg_host"
            echo -e "\t$reg_host/$reg_prefix/{{serviceName}}:{{lastBranchName}}"
            echo -e "\t$reg_host/$reg_prefix/{{serviceName}}:latest"
            echo ""
        done
    else
        echo "Error: registries file {{regFile}} not found."
        exit 1
    fi

# installs required tools
install-tools:
    #!/usr/bin/env bash
    if which go &> /dev/null; then
        echo "install goreleaser..."
        go install github.com/goreleaser/goreleaser@latest

        echo "install git-changelog..."
        go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest

        echo "install go testsum..."
        go install gotest.tools/gotestsum@latest

        echo "install golangci-lint"
        go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    fi