# arvan-challenge
arvan cloud code challenge

### commands
you can run following commands using `justfile` command runner
```shell
Available recipes:
    build                   # clean and build project
    buildimg tag="auto" cache="true" latest="true" # build service docker image
    buildimg-bin tag="auto" latest="true" # build service docker image from built binary
    clean                   # clean build directory
    default                 # list of available commands
    gen-changelog           # generate changelog file
    gorelease local="false" # release using goreleaser
    info                    # prints package informations
    install-tools           # installs required tools
    linter fix="false"      # run golang linter and fix it
    list-tags               # prints list of all tags
    publishimg tag="auto" latest="true" # publish builds and pushes docker image to registry
    pushimg tag="auto" latest="true" # push docker image to registry with tag
    registries              # prints list of all registries
    release target="patch" gorelease="false" publishimg="false" # release target=major/minor/patch
    rm-last-tag             # removes last tag from local and remote
    test                    # run go tests
    upx                     # build and compress binary
```