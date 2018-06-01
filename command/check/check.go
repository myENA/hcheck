package check

import (
	"fmt"

	// shared stuff
	"github.com/myENA/hcheck/shared"
)

// checkHost
func (c *Command) checkHost() error {
	var code int
	var err error

	// get return code and check error
	if code, err = shared.HTTPCode(
		c.config.URL.String(),
		c.config.Timeout,
		c.config.Insecure,
	); err != nil {
		return err
	}

	// validate response code
	if c.config.Code != code {
		return fmt.Errorf("expected %d, got %d",
			c.config.Code, code)
	}

	// all okay
	return nil
}
