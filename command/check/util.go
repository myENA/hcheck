package check

import (
	"flag"
	"fmt"
	"os"

	// shared stuff
	"github.com/myENA/hcheck/shared"
)

// setupFlags initializes the instance configuration
func (c *Command) setupFlags(args []string) error {
	var cmdFlags *flag.FlagSet // instance flagset
	var err error

	// init config if needed
	if c.config == nil {
		c.config = new(Config)
	}

	// init flagset
	cmdFlags = flag.NewFlagSet("check", flag.ContinueOnError)
	cmdFlags.Usage = func() { fmt.Fprint(os.Stdout, c.Help()); os.Exit(0) }

	// add shared config flags
	c.config.Config.Flags(cmdFlags)

	// parse flags and ignore error
	if err = cmdFlags.Parse(args); err != nil {
		return nil
	}

	// check for remaining garbage
	if cmdFlags.NArg() > 0 {
		return shared.ErrUnknownArg
	}

	// validate shared config and return
	return c.config.Config.Validate()
}
