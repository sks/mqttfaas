package outputprocessor

//Publisher ...
type Publisher interface {
	Publish(string, interface{}) error
}
