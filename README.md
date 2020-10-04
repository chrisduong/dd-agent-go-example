# Simple Go web with Dogstatsd

## Docker build

TODO: Makefile to create a task to build the Binary for Linux

## Origin Detection

You need to run with UDS protocol. SEE: [origin-detection](https://docs.datadoghq.com/developers/dogstatsd/unix_socket/?tab=host#origin-detection)

```ini
DD_DOGSTATSD_ORIGIN_DETECTION=true
DD_DOGSTATSD_TAG_CARDINALITY=high
DD_DOGSTATSD_SOCKET=/var/run/dd/dogstatd.sock
```

### low and orchestrator

Almost gives the same tags

### high

You've got container_name, container_id,..

