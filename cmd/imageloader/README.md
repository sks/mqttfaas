# Image loader

Exposes an API to load images to the docker container

## Building

```sh
docker build $DOCKER_BUILD_ARGS -t imageloader .
```

## Load Image

```sh
docker save -o image.tar.gz <docker_image_you_want_to_save>
http -f POST http://localhost:8000/load file@image.tar.gz
```