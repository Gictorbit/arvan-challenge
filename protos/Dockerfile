# syntax=docker/dockerfile:1.4
ARG GITHUB_ORIGIN=""
FROM swaggerapi/swagger-ui:latest

COPY docs/openapiv3/openapi.yaml /oapi3/openapi.yaml

ENV SWAGGER_JSON=/oapi3/openapi.yaml
ARG GITHUB_ORIGIN
LABEL org.opencontainers.image.source="${GITHUB_ORIGIN}"
EXPOSE 8080