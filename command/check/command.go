package check

import (
	"fmt"
	stdLog "log"

	// shared stuff
	"github.com/myENA/hcheck/shared"
)

// Config contains the command configuration
type Config struct {
	shared.Config
}

// Command is a Command implementation
type Command struct {
	Self   string
	Log    *stdLog.Logger
	config *Config
}

// Run is a function to run the command
func (c *Command) Run(args []string) int {
	var err error

	// init flags
	if err = c.setupFlags(args); err != nil {
		c.Log.Printf("[Error] Failed to init flags: %s", err.Error())
		return 1
	}

	// check host
	if err = c.checkHost(); err != nil {
		c.Log.Printf("[Error] Check failed: %s", err.Error())
		return 1
	}

	// exit clean
	return 0
}

// Synopsis shows the command summary
func (c *Command) Synopsis() string {
	return "Validate host return code and exit."
}

// Help shows the detailed command options
func (c *Command) Help() string {
	return fmt.Sprintf(`
Usage: %s cmd [options]

	Validate host return code and exit.
	Process will exit 0 on success and 1 on error.

Options:

	-url        The fully qualified URL to check
	-code       The expected return code (default: 200)
	-timeout    Time to wait for a response (default: 5s)
	-insecure   Skip TLS validation (default: false)
`, c.Self)
}
