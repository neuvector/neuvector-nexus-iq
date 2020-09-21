# NeuVector Nexus IQ integration

![Build](https://github.com/neuvector/neuvector-nexus-iq/workflows/Build/badge.svg)

The NeuVector Nexus IQ integration reports vulnerabilities detected by NeuVector to Sonatype Nexus IQ for reporting and analysis purposes.

For further project and research documentation refer to the contents of the sub folder `./docs`.

- [Guide: Setup integration between NeuVector and Nexus IQ](./docs/setup-guide.md)

## Features

- Listen to HTTP web hook calls triggered by registry scan events (`Registry.Scan.Report`) and runtime scan events (`Container.Scan.Report`) by the NeuVector controller
- Fetch registry and runtime scan report upon event occurrence from the NeuVector Controller API
    - Use GET `/v1/scan/registry/:name/image/:id` for registry scan reports
    - Use GET `/v1/scan/workload/:id` for runtime scan report
- Convert the NeuVector registry and runtime scan report contained in the payload of a web hook call to the CycloneDX format
- Identify the Nexus IQ application by image meta data
- Create a Nexus IQ application if an application cannot be inferred from image meta data
- Transmit the results to Nexus IQ via the Third-Party Scan REST API

## Requirements

The integration has been tested with the following versions:

- NeuVector Controller `>= 3.2.4`
- Nexus IQ `>= Release 94`

For versions not meeting these requirements, the function of the integration cannot be guaranteed.

## Docker image

Releases are available as Docker images [via DockerHub](https://hub.docker.com/r/neuvector/neuvector-nexus-iq). Images are tagged correspond to release versions.

```bash
# Latest release
docker pull neuvector/neuvector-nexus-iq:latest

# Specific release
docker pull neuvector/neuvector-nexus-iq:v1.0.0
```

## Build

Build the executable

```
make build
```

Execute unit tests

```
make test
```

## Development

Generate client and schema code

```
make generate
```

## Usage

### Native executable

```
./nv-nx-iq --help
./nv-nx-iq serve --help
```

```
./nv-nx-iq serve \
    --address 0.0.0.0 \
    --port 5080 \
    --nv-endpoint "https://127.0.0.1:10443" \
    --nv-username admin \
    --nv-password admin \
    --nx-endpoint "http://127.0.0.1:8070" \
    --nx-org "Sandbox Organization" \
    --nx-username admin \
    --nx-password admin123
```

For details on the configuration arguments, refer to the dedicated section on configuration below.

### Docker image

```
docker run --rm neuvector/neuvector-nexus-iq:latest --help
docker run --rm neuvector/neuvector-nexus-iq:latest serve --help
```

```
docker run --rm neuvector/neuvector-nexus-iq:latest serve \
    --address 0.0.0.0 \
    --port 5080 \
    --nv-endpoint "https://127.0.0.1:10443" \
    --nv-username admin \
    --nv-password admin \
    --nx-endpoint "http://127.0.0.1:8070" \
    --nx-org "Sandbox Organization" \
    --nx-username admin \
    --nx-password admin123
```

## Configuration

The integration can be configured via command line arguments, environment variables or via a YAML configuration file. The `--config` argument can reference a configuration file. An example YAML configuration file is available at `./test/config/example.yaml`.

The following table provides an overview of available configuration parameters.

Key | Argument | Environment variable | Description | Default | Example
--- | --- | --- | --- | --- | ---
`address` | `--address` | `NV_NX_ADDRESS` | Address of the webhook server | `127.0.0.1` | 
`port` | `--port` | `NV_NX_PORT` | Port of the webhook server | `5080` | 
`neuvector.endpoint` | `--nv-endpoint` | `NV_NX_NEUVECTOR_ENDPOINT` | Endpoint of the NeuVector Controller REST API |  | `https://127.0.0.1:10443`
`neuvector.username` | `--nv-username` | `NV_NX_NEUVECTOR_USERNAME` | Username of the NeuVector Controller | | `admin`
`neuvector.password` | `--nv-password` | `NV_NX_NEUVECTOR_PASSWORD` | Password of the NeuVector Controller | | `admin`
`neuvector.insecure` | `--nv-insecure` | `NV_NX_NEUVECTOR_INSECURE` | If set, TLS certificate verification is skipped for the NeuVector controller. This should be used in testing scenarios only. | `false` | 
`nexusiq.endpoint` | `--nx-endpoint` | `NV_NX_NEXUSIQ_ENDPOINT` | Endpoint of the Nexus IQ REST API | | `http://127.0.0.1:8070`
`nexusiq.username` | `--nx-username` | `NV_NX_NEXUSIQ_USERNAME` | Username of Nexus IQ | | `admin`
`nexusiq.password` | `--nx-password` | `NV_NX_NEXUSIQ_PASSWORD` | Password of Nexus IQ | | `admin123`
`nexusiq.insecure` | `--nx-insecure` | `NV_NX_NEXUSIQ_INSECURE` | If set, TLS certificate verification is skipped for NexusIQ. This should be used for testing scenarios only. | `false` | 
`nexusiq.source` | `--nx-source` | `NV_NX_NEXUSIQ_SOURCE` | Source of vulnerabilities in Nexus IQ report | `NeuVector` | 
`nexusiq.organization_name` | `--nx-org` | `NV_NX_NEXUSIQ_ORGANIZATION_NAME` | Name of the Nexus IQ organization to which vulnerabilities are reported |  | `Sandbox Organization`
`nexusiq.app_name_label` | `--nx-app-name-label` | `NV_NX_NEXUSIQ_APP_NAME_LABEL` | Key of the label from which the name of the Nexus IQ application is inferred. If not provided, the application name will be derived from the name of the image. | `"com.sonatype.nexus.iq.applicationName` | 

## Integration tests

### Test infrastructure

- Based on Docker Compose and require Docker on the workstation
- Verified using Docker Desktop 2.3.0.3+ on macOS Catalina

## Prerequisites

- License files are required for NeuVector and Nexus IQ and must be placed in `./test/licenses`
    - NeuVector license is required at `./test/licenses/neuvector.txt`
    - Nexus IQ license is required at `./test/licenses/nexusiq.lic`

### Start test infrastructure

```
make e2e-infrastructure-start
```

After starting the integration test infrastructure:

- the NeuVector Controller UI is available at https://127.0.0.1:8443 with username `admin` and password `admin`
- the NeuVector Controller REST API is available at https://127.0.0.1:10443 with username `admin` and password `admin`
- the Nexus IQ UI is available at http://127.0.0.1:8070 with username `admin` and password `admin123`
- a private Docker registry at http://127.0.0.1:5000 without authentication

Manual tasks after the first time starting the test infrastructure on a machine. These tasks are currently not automated as part of the integration tests.

- Configure the Nexus IQ license in Nexus IQ Web UI
- Add exactly one image to the private Docker registry
    - As part of the registry scan test all images of the private registry will be scanned

### Run integration tests

Run all integration tests:

```
make e2e-test
```

Run selected integration tests:

```
make e2e-test TEST=TestRegistryScan
make e2e-test TEST=TestContainerScan
```
