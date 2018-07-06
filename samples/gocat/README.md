# Catter function

```sh
docker build -t gocatter:latest .

echo "asdasdadasda" | docker run -e "FIRED_BY=FIRED_BY_MQTT_TOPIC" --rm -i gocatter:latest
```