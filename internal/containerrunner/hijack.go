package containerrunner

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stdcopy"
)

// A hijackedIOStreamer handles copying input to and output from streams to the
// connection.
type hijackedIOStreamer struct {
	inputStream  io.ReadCloser
	outputStream io.Writer
	errorStream  io.Writer
	resp         types.HijackedResponse
}

// stream handles setting up the IO and then begins streaming stdin/stdout
// to/from the hijacked connection, blocking until it is either done reading
// output, the user inputs the detach key sequence when in TTY mode, or when
// the given context is cancelled.
func (h *hijackedIOStreamer) stream(ctx context.Context) error {

	outputDone := h.beginOutputStream()
	inputErr := h.beginInputStream()

	select {
	case err := <-outputDone:
		return err
	case err := <-inputErr:
		if err != nil {
			return err
		}
		select {
		case err := <-outputDone:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (h *hijackedIOStreamer) beginOutputStream() <-chan error {
	outputDone := make(chan error)
	go func() {
		_, err := stdcopy.StdCopy(h.outputStream, h.errorStream, h.resp.Reader)

		outputDone <- err
	}()

	return outputDone
}

func (h *hijackedIOStreamer) beginInputStream() (doneC <-chan error) {
	inputErr := make(chan error)

	go func() {
		_, err := io.Copy(h.resp.Conn, h.inputStream)

		if err != nil {
			inputErr <- err
		}

		if err := h.resp.CloseWrite(); err != nil {
			inputErr <- err
		}

		close(inputErr)
	}()

	return inputErr
}
