# Guide: Setup integration between NeuVector and Nexus IQ

This guide describes how to setup the integration between NeuVector and Nexus IQ using the NeuVector Nexus IQ integration (referred to as *nv-iq-integration* in the following).

## Prerequisites

### NeuVector controller

- A NeuVector controller with a valid license applied is required
- This guide assumes the [REST API endpoint of the controller](https://docs.neuvector.com/basics/installation/console#connect-to-rest-api-server) to be `https://127.0.0.1:10443`
- It is recommended to setup a [dedicated NeuVector controller user](https://docs.neuvector.com/basics/users) for the nv-iq-integration
- This guide assumes a user `nv-nx-iq` with password `qi-nx-vn` is configured in the NeuVector controller 

### Nexus IQ server

- A Nexus IQ server with a valid license applied is required
- This guide assumes the [REST API endpoint of the IQ server](https://help.sonatype.com/iqserver/automating/rest-apis) to be `http://127.0.0.1:8070`
- It is recommended to setup a [dedicated NeuVector controller user](https://docs.neuvector.com/basics/users) for the nv-iq-integration
- This guide assumes a user `nv-nx-iq` with password `qi-nx-vn` is configured in the Nexus IQ server

## Overview

- Webhooks for registry scan events and runtime scan events are enabled in the NeuVector controller
- The nv-iq-integration receives webhook calls from the NeuVector controller and creates third-party scan reports in the Nexus IQ server
- The nv-iq-integration runs as a separate service in a Docker container

## Steps

### Configure webhook in NeuVector controller

The endpoint of the nv-iq-integration service is configured as the webhook receiver in the NeuVector controller.

- In the NeuVector UI, open *Settings* - *Configuration*
- Set the *Webhook Url* to the webhook endpoint of the nv-iq-integration service
    - Example: `http://localhost:5080/webhook`
    - When running the nv-iq-integration service as a container in an evaluation setup with Docker Desktop, the *Webhook Url* is `http://host.docker.internal:5080/webhook`

### Configure response rules in NeuVector controller

[Response rules](https://docs.neuvector.com/policy/responserules) are configured in the NeuVector controller, so that registry scan events and runtime scan events are reported to the webhook receiver.

- In the NeuVector UI, open *Policy* - *Response rules*

To report on *registry scan events*:

- Add a new rule
- Category: `CVE-Report`
- Group is optional
- Criteria: `name:Registry.Scan.Report`
- Action: Webhook
- Status: Enabled

To report on *runtime scan events*:

- Add a new rule
- Category: `CVE-Report`
- Group is optional
- Criteria: `name:Container.Scan.Report`
- Action: Webhook
- Status: Enabled

> The nv-iq-integration service will only receive webhooks when the response rules are enabled.

- It is recommended to create separate response rules for registry and container scan events.
- Optionally, additional criteria can be added to the response rules

### Start the nv-iq-integration service

The following command starts the nv-iq-integration service on a local Docker (Desktop) host and pulls the image from the [official DockerHub repository neuvector/neuvector-nexus-iq](https://hub.docker.com/r/neuvector/neuvector-nexus-iq).

```bash
docker run -d --rm -p 5080:5080 neuvector/neuvector-nexus-iq:latest serve \
    --address 0.0.0.0 \
    --port 5080 \
    --nv-endpoint "https://127.0.0.1:10443" \
    --nv-username nv-nx-iq \
    --nv-password qi-nx-vn \
    --nx-endpoint "http://127.0.0.1:8070" \
    --nx-org "Sandbox Organization" \
    --nx-username nv-nx-iq \
    --nx-password qi-nx-vn
```

As an alternative to the arguments, the nv-iq-integration service can be configured using a YAML config file.

```bash
docker run -d --rm -p 5080:5080 neuvector/neuvector-nexus-iq:latest serve \
    --config ./neuvector-nexus-iq.yaml
```

Example YAML config file:

```yaml
address: 127.0.0.1
port: 5080

neuvector:
  endpoint: https://127.0.0.1:10443
  insecure: false
  username: nv-nx-iq
  password: qi-nx-vn

nexusiq:
  endpoint: https://127.0.0.1:8070
  insecure: false
  username: nv-nx-iq
  password: qi-nx-vn
  organization_name: Sandbox Organization
```

Configuration can also be sourced from environment variables. All command line arguments can be provided by and environment variable with the with `NV_NX_` followed by the uppercase argument. For example, the `--address` argument can be set by the environment variable `NV_NX_ADDRESS`.

### Trigger a registry scan event

To test the integration, trigger a registry scan event by starting a scan of a [registry configured in the NeuVector controller](https://docs.neuvector.com/scanning/registry).

This can be accomplished in the NeuVector UI in the *Assets* - *Registries* view via the button *Start scan* with a registry selected.

### Trigger a container scan event

To test the integration with a container scan event, start a container workload on one of the hosts which is covered by the NeuVector installation.

In an evaluation environment based on Docker Desktop, start a container.
