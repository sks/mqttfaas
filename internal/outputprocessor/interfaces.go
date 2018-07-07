package outputprocessor

//Publisher ...
//go:generate counterfeiter . Publisher
type Publisher interface {
	Publish(string, interface{}) error
}
