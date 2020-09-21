PROJECT_PATH := $(abspath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

GO_MODULE := github.com/neuvector/neuvector-nexus-iq

CLI_EXECUTABLE_NAME := nv-nx-iq

# Image

IMAGE_NAME := neuvector/neuvector-nexus-iq

IMAGE_TAG := latest

# Build variables

BUILD_VERSION = $(shell git describe --tags --always)

BUILD_COMMIT = $(shell git rev-parse HEAD)

BUILD_TIME = $(shell date -u)

# # CycloneDX schema

# ## CycloneDX XSD root schema
pkg/cyclonedx/schema/cyclonedx-1.1.xsd:
	mkdir -p $(dir $@)
	curl --silent -o $@ https://cyclonedx.org/schema/bom/1.1

# ## CycloneDX XSD spdx schema
pkg/cyclonedx/schema/cyclonedx-spdx-1.0.xsd:
	mkdir -p $(dir $@)
	curl --silent -o $@ https://cyclonedx.org/schema/spdx

# ## CycloneDX XSD vulnerabilities schema
pkg/cyclonedx/schema/cyclonedx-vulnerability-1.0.xsd:
	mkdir -p $(dir $@)
	curl --silent -o $@ https://cyclonedx.org/schema/ext/vulnerability/1.0

# # Test licenses

.PHONY: licenses
licenses: \
	test/licenses/neuvector.txt \
	test/licenses/nexusiq.lic \

test/licenses/neuvector.txt:
	$(error NeuVector license must be available at $@)

test/licenses/nexusiq.lic:
	$(error Nexus IQ license must be available at $@)

# # Integration test infrastructure

test/e2e/docker-compose.yml: \
	test/neuvector-controller/docker-compose.yml \
	test/nexus-iq/docker-compose.yml \
	test/docker-registry/docker-compose.yml
	yq m $^ > $@

test/e2e/docker-registry-config.yml: \
	test/docker-registry/docker-registry-config.yml
	cp $< $@ 

# e2e-infrastructure

.PHONY: e2e-infrastructure-start
e2e-infrastructure-start: \
	test/e2e/docker-compose.yml \
	test/e2e/docker-registry-config.yml
	cd test/e2e && docker-compose up --no-recreate

.PHONY: e2e-infrastructure-stop
e2e-infrastructure-stop:
	cd test/e2e && docker-compose down

.PHONY: e2e-infrastructure-status
e2e-infrastructure-status:
	cd test/e2e && docker-compose ps

# # CLI executable

$(CLI_EXECUTABLE_NAME): .FORCE
	go build -o $@ -ldflags="-X '$(GO_MODULE)/build.Version=$(BUILD_VERSION)' -X '$(GO_MODULE)/build.Commit=$(BUILD_COMMIT)' -X '$(GO_MODULE)/build.Time=$(BUILD_TIME)'" .
	chmod +x $@

# # Short targets

.FORCE:

.PHONY: generate
generate:
	go generate ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build: $(CLI_EXECUTABLE_NAME)

# Execute unit tests
.PHONY: test
test:
	go test -v ./...

# Execute integration tests
# `-count=1`: Disable test result cache
# `-v`: Verbose output
.PHONY: e2e-test
e2e-test: licenses
	go test -v -count 1 -tags e2e$(if $(TEST), -run $(TEST),) ./e2e

.PHONY: image
image:
	DOCKER_BUILDKIT=1 docker build \
	--build-arg 'BUILD_VERSION=$(BUILD_VERSION)' \
	--build-arg 'BUILD_COMMIT=$(BUILD_COMMIT)' \
	--build-arg 'BUILD_TIME=$(BUILD_TIME)' \
	-t $(IMAGE_NAME):$(IMAGE_TAG) \
	.

.PHONY: container-version
container-version:
	docker run --rm $(IMAGE_NAME):$(IMAGE_TAG) version