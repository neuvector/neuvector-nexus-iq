# NeuVector Nexus IQ integration

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

## Build

Generate client and schema code

```
make generate
```

Build the executable

```
make build
```

Execute unit tests

```
make test
```

## Usage

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

`nv-nx-iq` can optionally be configured using a YAML configuration file referenced in the `--config` argument. An example YAML configuration file is available at `./test/config/example.yaml`.

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
- the Nexus IQ UI is available at https://127.0.0.1:8070 with username `admin` and password `admin123`
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
