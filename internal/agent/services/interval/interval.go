package interval

import (
	"time"
)

func StartInterval(intervalTime time.Duration, runner func()) chan bool {
	ticker := time.NewTicker(intervalTime)
	// quit unused
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				runner()
			case <-quit:
				println("hi")
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}
