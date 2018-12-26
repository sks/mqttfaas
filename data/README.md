# Running

```sh
# To use openfaas as the provider
docker-compose \
    -f openfaas.yml \
    -f httpinvoker.yml \
    -f databus.yml up

docker pull functions/alpine

docker save -o image.tar.gz functions/alpine
http -f POST http://localhost:8000/load file@image.tar.gz

faas-cli deploy \
    --image functions/alpine \
    --name wc \
    --network faas-swarm_default \
    --fprocess wc
# Subscribe
mqttcli sub -t "/dev/+" -dd

# Publish message
mqttcli pub -t "/wordcount/input" -m "a message should be posted in /dev/stdout"
```

## Using fn

```sh
docker-compose \
    -f fn.yml \
    -f httpinvoker.yml \
    -f databus.yml up
```