package watch

import (
	"context"
	"fmt"
	stdLog "log"
	"sync"
	"time"

	// shared stuff
	"github.com/myENA/hcheck/shared"
)

// Config contains the command configuration
type Config struct {
	shared.Config
	interval time.Duration
}

// Command is a Command implementation
type Command struct {
	Self   string
	Log    *stdLog.Logger
	config *Config
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
}

// Run is a function to run the command
func (c *Command) Run(args []string) int {
	var err error

	// init flags
	if err = c.setupFlags(args); err != nil {
		c.Log.Printf("[Error] Failed to init flags: %s", err.Error())
		return 1
	}

	// init waitgroup
	c.wg = new(sync.WaitGroup)

	// init context
	c.ctx, c.cancel = context.WithCancel(context.Background())

	// start watcher
	go c.watchHost()

	// block and wait
	return c.sigWait()
}

// Synopsis shows the command summary
func (c *Command) Synopsis() string {
	return "Poll target host and validate return code."
}

// Help shows the detailed command options
func (c *Command) Help() string {
	return fmt.Sprintf(`
Usage: %s cmd [options]

	Poll target host and validate return code.
	Process will run until interrupted.

Options:

	-url        The fully qualified URL to check
	-code       The expected return code (default: 200)
	-timeout    Time to wait for a response (default: 5s)
	-insecure   Skip TLS validation (default: false)
	-interval   Watch poll interval (default: 5s)
`, c.Self)
}
