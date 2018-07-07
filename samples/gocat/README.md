# Catter function

This functions echoes the message that was posted

## Usage

```sh
docker build -t sabithksme/mqttfaas_gocat .

echo "asdasdadasda" | docker run -e "FIRED_BY=FIRED_BY_MQTT_TOPIC" --rm -i sabithksme/mqttfaas_gocat:latest
```