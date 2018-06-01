package main

import (
	stdLog "log"
	"os"

	// cli library
	"github.com/mitchellh/cli"

	// our command implementations
	"github.com/myENA/hcheck/command/check"
	"github.com/myENA/hcheck/command/watch"
)

// package global logger
var logger *stdLog.Logger

// available commands
var cliCommands map[string]cli.CommandFactory

// init command factory
func init() {
	// init logger
	logger = stdLog.New(os.Stderr, "", stdLog.LstdFlags)

	// register sub commands
	cliCommands = map[string]cli.CommandFactory{
		"check": func() (cli.Command, error) {
			return &check.Command{
				Self: os.Args[0],
				Log:  logger,
			}, nil
		},
		"watch": func() (cli.Command, error) {
			return &watch.Command{
				Self: os.Args[0],
				Log:  logger,
			}, nil
		},
	}
}
