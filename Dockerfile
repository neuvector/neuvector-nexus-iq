# syntax = docker/dockerfile:experimental

# Base image tag for build stage
# https://hub.docker.com/_/golang?tab=tags&page=1&name=alpine
ARG GOLANG_IMAGE_TAG=1.14.4-alpine3.12

# # Base
FROM golang:${GOLANG_IMAGE_TAG} as base

# # Development environment
FROM base as dev

# ## Packages
RUN apk add -U --no-cache \
    ca-certificates \
    bash \
    make

# ## Work directory
RUN mkdir -p /app
WORKDIR /app

# ## Shell entrypoint
ENTRYPOINT ["/bin/bash"]

# # Build
FROM dev as build

# ## Build variables
ARG GO_MODULE=github.com/neuvector/neuvector-nexus-iq

ARG BUILD_VERSION=dev
ARG BUILD_COMMIT=
ARG BUILD_TIME=

# ## Add dependencies
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

# ## Download dependencies
RUN --mount=type=cache,id=nv-nx-iq-go-pkg,target=/go/pkg go mod download

# ## Add sources
COPY . /app

# ## Build executable
RUN --mount=type=cache,id=nv-nx-iq-go-pkg,target=/go/pkg CGO_ENABLED=0 go build \
    -o /app/nv-nx-iq \
    -ldflags="-X '${GO_MODULE}/build.Version=${BUILD_VERSION}' -X '${GO_MODULE}/build.Commit=${BUILD_COMMIT}' -X '${GO_MODULE}/build.Time=${BUILD_TIME}'"

# # Test
FROM build as test

# Run unit tests (excluding integration tests)
RUN --mount=type=cache,id=nv-nx-iq-go-pkg,target=/go/pkg CGO_ENABLED=0 go test -count 1 ./...

# # Application
FROM scratch as app

USER 1000

COPY --from=test /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=test /app/nv-nx-iq /nv-nx-iq

ENTRYPOINT ["/nv-nx-iq"]
CMD ["serve"]
