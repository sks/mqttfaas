# Catter function

```sh
docker build -t sabithksme/mqttfaas_gocat .

echo "asdasdadasda" | docker run -e "FIRED_BY=FIRED_BY_MQTT_TOPIC" --rm -i gocatter:latest
```