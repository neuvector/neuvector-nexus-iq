# syntax = docker/dockerfile:experimental

# Base image tag for build stage
# https://hub.docker.com/_/golang?tab=tags&page=1&name=alpine
ARG GOLANG_IMAGE_TAG=1.14.4-alpine3.12

ARG APP_NAME=nv-nx-iq

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

# ## Third-party development dependencies
# TODO swagger
# TODO shell

# # Build
FROM dev as build

# ## Add dependencies
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

# ## Download dependencies
RUN --mount=type=cache,id=nv-nx-iq-go-pkg,target=/go/pkg go mod download

# ## Add sources
COPY . /app

# ## Build executable
RUN --mount=type=cache,id=nv-nx-iq-go-pkg,target=/go/pkg CGO_ENABLED=0 go build -o /app/nv-nx-iq

# # Test
FROM build as test

# Run unit tests
#RUN --mount=type=cache,id=nv-nx-iq-go-pkg,target=/go/pkg CGO_ENABLED=0 go test ./...

# # Application
FROM scratch as app

USER 1000

COPY --from=test /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=test /app/nv-nx-iq /nv-nx-iq

ENTRYPOINT ["/nv-nx-iq"]
CMD ["serve"]
