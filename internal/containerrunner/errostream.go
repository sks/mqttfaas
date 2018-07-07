package containerrunner

import (
	"log"
	"os"
)

var l = log.New(os.Stderr, "", log.LstdFlags)

type errorStream struct {
	identifier string
}

func newErrorStream(identifier string) *errorStream {
	return &errorStream{
		identifier,
	}
}

func (e *errorStream) Write(d []byte) (int, error) {
	l.Printf("function=%s:\t%s\n", e.identifier, string(d))
	return len(d), nil
}
