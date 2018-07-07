# Temperature Converter

This function container listens to the topic `house/+/room/+/temperature` as dictated by [Dockerfile](./Dockerfile)

## Usage

### Testing

```sh
docker build -t sabithksme/mqttfaas_temp_converter .

echo "39" | docker run -e "FIRED_BY=house/388/room/living/temperature" \
    --rm -i sabithksme/mqttfaas_temp_converter
```

### Usage  with MQTT FAAS

```sh
docker pull sabithksme/mqttfaas_temp_converter

export MQTT_HOST="localhost"

mqttcli sub -t "house/+/room/+/#"

mqttcli pub -t "house/388/room/living/temperature" -m "39"

# You should see an output of 102.2 on the subscription channel
```