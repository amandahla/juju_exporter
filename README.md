# Juju Prometheus Exporter

Simple Juju Prometheus exporter written in Golang. Exports metrics regarding the current status of the Juju agents for all machines, applications and units of a model. Uses the [Go client](https://github.com/juju/juju/tree/develop/api) to connect to Juju, so the Juju binaries are not required.

Juju exporter supports querying metrics for multiple Juju models with a single config file.

## Usage

- [Download binaries for latest release](https://github.com/neoaggelos/juju_exporter/releases)
- Run with Docker: `docker run --rm -it -v $PWD/juju_exporter.yaml:/juju_exporter.yaml -p 9970:9970 neoaggelos/juju_exporter`
- Snap (TODO)

## Releases

`juju_exporter` uses [GitHub](https://github.com/neoaggelos/juju_exporter/releases) for releases.

## Setup

**Recommended**: Create a Juju user named `juju-exporter` with read-only access to the model you want to export metrics from. With the `juju` CLI:

```bash
$ juju add-user juju-exporter
$ juju change-user-password juju-exporter
$ juju grant juju-exporter read $MODEL_NAME
```

Create a `juju_exporter.yaml` configuration file with the following contents (You can find most that information for your own deployment by running `juju show-controller`).

```yaml
---
default: MODEL_1
models:
  MODEL_1:
    api-endpoints: [10.0.0.1:17070, 10.0.0.2:17070, 10.0.0.3:17070]
    model-uuid: MODEL_UUID_1
    username: juju-exporter
    password: super-safe-password
    ca-cert: |
      -----BEGIN CERTIFICATE-----
      <certificate contents>
      -----END CERTIFICATE-----
  MODEL_2:
    api-endpoints: [10.0.0.1:17070, 10.0.0.2:17070, 10.0.0.3:17070]
    model-uuid: MODEL_UUID_2
    username: juju-exporter
    password: super-safe-password
    ca-cert: |
      -----BEGIN CERTIFICATE-----
      <certificate contents>
      -----END CERTIFICATE-----
```

## Query

Query metrics for `MODEL_2`:

- `curl http://localhost:9970/metrics?model=MODEL_2`

Or query for the default model:

- `curl http://localhost:9970/metrics`
