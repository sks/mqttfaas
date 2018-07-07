# Using Docker in Docker

At times you dont want to use the docker from host machine.

We have [DIND images](https://hub.docker.com/_/docker/) thanks to docker.

```sh
docker-compose \
    -f docker-compose.yml \
    -f integration/dind/docker-compose.yml up \
    -d

export DOCKER_HOST=tcp://localhost:12375
docker pull sabithksme/mqttfaas_gocat

# Ready to roll by pushing messages to MQTT and see the functions get invoked
```