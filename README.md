# HTTP Invoker

[![Build Status](https://travis-ci.org/sks/mqttfaas.svg?branch=master)](https://travis-ci.org/sks/mqttfaas) [![Go Report Card](https://goreportcard.com/badge/github.com/sks/mqttfaas)](https://goreportcard.com/report/github.com/sks/mqttfaas)

Call HTTP API's based on MQTT Messages.

[httpinvoker](./cmd/httpinvoker) looks for [config.yml](./data/mounts/httpinvoker/data/config.yml) to figure out which http api to call to make based on mqtt topic

## Why

- [ ] Developers dont have to worry about the messaging framework.
- [ ] Developers concentrate on business function
- [ ] Reuse functions

## Getting Started

### Using [openfaas](https://github.com/openfaas/faas)

A sample [docker-compose.yml](./data/openfaas.yml) entry is provided.

```sh
cd data

docker-compose \
    -f openfaas.yml \
    -f httpinvoker.yml \
    -f databus.yml up

```

### Using [fn](http://fnproject.io)

A sample [docker-compose.yml](./data/fn.yml) entry is provided.

```sh
cd data

docker-compose \
    -f fn.yml \
    -f httpinvoker.yml \
    -f databus.yml up
```