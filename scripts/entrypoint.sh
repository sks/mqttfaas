#!/bin/sh -e

if [ $@ -ne 0 ]
then
    dockerd-entrypoint.sh $@
    exit 0
fi
dockerd-entrypoint.sh $@ &

sleep 2

/bin/mqttfaas