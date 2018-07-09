# Function Metadata

## Available labels for each functions

- `mqtt_faas` Registers as a function that can be invoked on MQTT Topic
    [Required] for a image to be seen by MQTT-FAAS to picked up

- `mqtt_faas_topic` Topics on which the functions should be invoked
    [Optional] If not set then the function is invoked on all the message on all the topics.

    The values can be in the format of a [valid MQTT Topic](https://www.hivemq.com/blog/mqtt-essentials-part-5-mqtt-topics-best-practices)

- `mqtt_faas_single_use_only` Dont use hot containers. No Reuse
    [Optional] Marks the container for deletion after each invocation.

    If the function is for single use, the name of the container is choosen at random

- `mqtt_faas_no_fired_by` then the `FIRED_BY` environment variable is not available to containers
    [Optional] Handy if the function dont really care about why the function was triggered.

    All information required for the function should be contained in the MQTT Message.

    If the function has no interest in topic, the name of the container would have the format `mqttfaas_-function_image_name`. Otherwise the container name looks like `mqttfaas_topic_subscribed-function_image_name`. 

    Checkout the [tests](./imagerunnerinput_test.go) for more details

## Persistent storage

A `/data/` directory is mounted to the containers to store any persistent data.