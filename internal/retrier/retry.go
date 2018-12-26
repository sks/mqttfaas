package retrier

import "time"

//Call retries the function for the given attempts with sleep in between
func Call(attempts int, sleep time.Duration, callable func() error) (err error) {
	for i := 0; i < attempts; i++ {
		err = callable()
		if err == nil {
			return
		}
		time.Sleep(sleep)
	}
	return err
}
