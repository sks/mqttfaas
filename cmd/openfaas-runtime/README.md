# Image loader

Exposes an API to load images to the docker container

## Building

```sh

docker save -o faasswarm.tar.gz openfaas/faas-swarm:0.5.0

docker build $DOCKER_BUILD_ARGS -t swarm_init .
```

## Running

```sh
    docker run --rm --name swarm -p 8080:8080 --privileged swarm_init
```
