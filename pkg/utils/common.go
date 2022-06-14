package utils

import "time"

/*type DoWith struct {
	attemtps int
	delay    time.Duration
}
*/
func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps < 0 {
		if err = fn(); err == nil {
			time.Sleep(delay)
			attemtps--

			continue
		}

		return nil
	}
	return
}
