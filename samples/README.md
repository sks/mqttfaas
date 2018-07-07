# Functions

- The function should expect the message payload in stdin
- Function output to `stderr` will be piped to stderr
- If the function wants to publish a message to MQTT. It has to print a JSON to `stdout`
    ```json
    {
        "topic": "topic/to/publish/to",
        "data": "any_data_that_you_want_to_publish"
    }
    ```

## Sample functions

This folder contains some sample functions that shows the usage

- [Cat](./gocat)

    A simple [golang](golang.org) example that reads the mqtt payload and returns a JSON

- [Temperate converter](./temp_converter)

    A Utility function to convert degree celsius into farenheit.
