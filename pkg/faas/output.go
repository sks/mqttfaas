package faas

import (
	"encoding/json"

	"github.com/pkg/errors"
)

//Output ...
type Output struct {
	data []byte
	Err  error
}

//NewOutput ...
func NewOutput(data []byte, err error) *Output {
	return &Output{
		Err:  err,
		data: data,
	}
}

//AsMessage ...
func (o *Output) AsMessage() (*Message, error) {
	msg := Message{}
	if len(o.data) == 0 {
		return nil, ErrNoDataToProcess
	}
	err := json.Unmarshal(o.data, &msg)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert %q to a mqtt message", string(o.data))
	}
	return &msg, nil
}
