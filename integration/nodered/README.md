# Node red

This is mainly for testing purpose and easier way to debug MQTT.

Sample flows are checked in [here](../../data/nodered/flows.json)

```sh
docker-compose \
    -f docker-compose.yml \
    -f integration/nodered/docker-compose.yml up
```