package watch

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	cmdFlags = flag.NewFlagSet("watch", flag.ContinueOnError)
	cmdFlags.Usage = func() { fmt.Fprint(os.Stdout, c.Help()); os.Exit(0) }

	// add package flags
	cmdFlags.DurationVar(&c.config.interval, "interval", 5*time.Second,
		"watch interval")

	// add shared flags
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

// trap signals and wait
func (c *Command) sigWait() int {
	var sigChan = make(chan os.Signal, 1) // signal channel
	var sig os.Signal                     // captured signal

	// trap signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// handle signals
	for {
		select {
		case sig = <-sigChan:
			// log signal
			c.Log.Printf("[Info] Exiting on signal: %d (%s)",
				sig, sig)
			// cancel watcher context
			c.cancel()
			// wait for watcher to finish
			c.wg.Wait()
			// all done
			c.Log.Print("[Info] Watcher finished.")
			// exit clean
			return 0
		}
	}
}
