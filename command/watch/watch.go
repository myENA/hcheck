package watch

import (
	"time"

	// shared stuff
	"github.com/myENA/hcheck/shared"
)

// checkHost
func (c *Command) watchHost() {
	var code int
	var ticker *time.Ticker
	var err error

	// handle waitgroup
	c.wg.Add(1)
	defer c.wg.Done()

	// init ticker
	ticker = time.NewTicker(c.config.interval)

	for {
		select {
		case <-c.ctx.Done():
			c.Log.Print("[Info] Watcher exiting on closed context.")
			ticker.Stop()
			return
		case <-ticker.C:
			// get return code and check error
			if code, err = shared.HTTPCode(
				c.config.URL.String(),
				c.config.Timeout,
				c.config.Insecure,
			); err != nil {
				c.Log.Printf("[Error] Failed to get HTTP code: %s",
					err.Error())
				continue
			}

			// validate response code
			if c.config.Code != code {
				c.Log.Printf("[Error] Unexpexted return code: wanted %d, got %d",
					c.config.Code, code)
				continue
			}
		}
	}
}
