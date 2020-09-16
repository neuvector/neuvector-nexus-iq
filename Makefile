PROJECT_PATH := $(abspath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

GO_MODULE := github.com/neuvector/neuvector-nexus-iq

CMD_EXECUTABLE_NAME := nv-nx-iq

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

# # Cmd executable

$(CMD_EXECUTABLE_NAME): .FORCE
	go build -o $@ .
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
build: $(CMD_EXECUTABLE_NAME)

# Execute unit tests
.PHONY: test
test:
	go test -v ./...

# Execute integration tests
# -count=1 to disable test result cache
# -v verbose output
.PHONY: e2e-test
e2e-test: licenses
	go test -v -count 1 -tags e2e$(if $(TEST), -run $(TEST),) $(GO_MODULE)/e2e

.PHONY: image
image:
	DOCKER_BUILDKIT=1 docker build -t nv-nx-iq:latest .

