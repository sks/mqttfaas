# MQTT FAAS

[![Build Status](https://travis-ci.org/sks/mqttfaas.svg?branch=master)](https://travis-ci.org/sks/mqttfaas) [![Go Report Card](https://goreportcard.com/badge/github.com/sks/mqttfaas)](https://goreportcard.com/report/github.com/sks/mqttfaas)

Execute docker functions that are run based on mqtt messages.

## Quick Setup

```sh

# Download latest docker-compose.yml
wget https://raw.githubusercontent.com/sks/mqttfaas/master/docker-compose.yml

# Start docker container
docker-compose up
```

## Definition of functions

Samples can be found in [samples](./samples) folder

- Must have a LABEL `mqtt_faas`
- Topics of interest for each function should be marked using LABEL `mqtt_faas_topic`
- If the `mqtt_faas_topic` is empty. the function gets all the messages from all topics
- Topic messages are available to the functions on stdin
- Topic name is available as environment variable `FIRED_BY`

## Development environment

```sh
git clone github.com/sks/mqttfaas $GOPATH/src/github.com/sks/mqttfaas
cd $_

# Install Dependency and build binaries
make

# Start the MQTT
docker-compose up -d
```