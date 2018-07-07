# MQTT FAAS

[![Build Status](https://travis-ci.org/sks/mqttfaas.svg?branch=master)](https://travis-ci.org/sks/mqttfaas) [![Go Report Card](https://goreportcard.com/badge/github.com/sks/mqttfaas)](https://goreportcard.com/report/github.com/sks/mqttfaas)

Execute docker functions that are run based on mqtt messages.

## Definition of functions

1. Must have a LABEL `mqtt_faas`
2. Topics of interest for each function should be marked using LABEL `mqtt_faas_topic`
3. If the `mqtt_faas_topic` is empty. the function gets all the messages from all topics
4. Topic messages are available to the functions on stdin
5. Topic name is available as environment variable `FIRED_BY`

## Development environment

```sh
    git clone github.com/sks/mqttfaas $GOPATH/src/github.com/sks/mqttfaas
    cd $_

    # Start the MQTT
    docker-compose up -d

    go run cmd/mqttfaas/main.go

```

## Helpful make commands

```sh
    make dockerize
    make help

# # Creating mqtt faas that comes with dind
# make dockerize/dind

# docker run $DOCKER_RUN_ARGS \
#     --privileged \
#     -e "MQTT_CONNECTION_STRING=tcp://mqtthost:1883" \
#     --add-host="mqtthost:$(ifconfig | grep inet | grep "192.168"  | cut -d ' ' -f2)" --rm \
#     --name mqttfaas \
#     mqttfaas:dind


# docker run --rm --link mqttfaas:docker docker:dind docker-entrypoint.sh version
```