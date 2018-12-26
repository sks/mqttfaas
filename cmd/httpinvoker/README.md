# FN Invoker

Invoke the serverless function based on mqtt messages

## Building

```sh
docker build $DOCKER_BUILD_ARGS -t httpinvoker .
```

## Testing

```sh
# Run MQTT
docker inspect mqtt || docker run -d --name=mqtt -p 1883:1883 --rm eclipse-mosquitto

go run main.go

mqttcli pub -t "/wordcount/input" -m "a message should be posted in /dev/null"

mqttcli sub -t "#"
```