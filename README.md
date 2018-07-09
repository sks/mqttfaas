# MQTT FAAS

[![Build Status](https://travis-ci.org/sks/mqttfaas.svg?branch=master)](https://travis-ci.org/sks/mqttfaas) [![Go Report Card](https://goreportcard.com/badge/github.com/sks/mqttfaas)](https://goreportcard.com/report/github.com/sks/mqttfaas)

Execute functions that are run based on mqtt messages.

## Quick Setup

### Install the binary

```sh
curl https://raw.githubusercontent.com/sks/mqttfaas/master/install.sh | sh
```

### On Docker

```sh

# Download latest docker-compose.yml
wget https://raw.githubusercontent.com/sks/mqttfaas/master/docker-compose.yml

# Download a sample Function.
# Samples can be found in samples directory
docker pull sabithksme/mqttfaas_gocat

# Start docker container
docker-compose up -d

# Grab a mqtt cli for testing purpose.
# I am using https://github.com/shirou/mqttcli
# go get github.com/shirou/mqttcli

export MQTT_HOST="localhost"

mqttcli sub -t "cat/#"

mqttcli pub -t "cat/input/message" -m "this message should be echoed back to /cat/output"
```

## Definition of functions

Samples can be found in [samples](./samples) folder

Guideline on [labels](https://docs.docker.com/config/labels-custom-metadata/) can be found in [here](./pkg/types/README.md)

## Development environment

Checkout [Makefile](./Makefile) for all available commands

```sh
git clone github.com/sks/mqttfaas $GOPATH/src/github.com/sks/mqttfaas
cd $_

# Install Dependency and build binaries
make

make help
```
