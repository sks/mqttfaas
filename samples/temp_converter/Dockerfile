FROM busybox

LABEL mqtt_faas="temp_coverter" mqtt_faas_topic="house/+/room/+/temperature"

COPY convert.sh /bin/convert.sh

ENTRYPOINT [ "/bin/convert.sh" ]