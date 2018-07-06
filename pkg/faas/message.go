package faas

//Message ...
type Message struct {
	Topic string      `json:"topic"`
	Data  interface{} `json:"data"`
}
