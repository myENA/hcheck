package shared

import (
	"flag"
	"net/url"
	"time"
)

// Config represents shared config
type Config struct {
	urlString string
	URL       *url.URL
	Code      int
	Timeout   time.Duration
	Insecure  bool
}

// Flags adds flags to populate shared config
func (c *Config) Flags(cmdFlags *flag.FlagSet) {
	// shared flags
	cmdFlags.StringVar(&c.urlString, "url", "",
		"check url")
	cmdFlags.IntVar(&c.Code, "code", 200,
		"expected return code")
	cmdFlags.DurationVar(&c.Timeout, "timeout", time.Second*5,
		"request timeout")
	cmdFlags.BoolVar(&c.Insecure, "insecure", false,
		"skip tls verification")
}

// Validate validates shared config
func (c *Config) Validate() error {
	var err error

	// check url
	if c.urlString == "" {
		return ErrMissingURL
	}

	// parse/validate url
	if c.URL, err = url.Parse(c.urlString); err != nil {
		return err
	}

	// ensure valid scheme
	if c.URL.Scheme == "" {
		c.URL.Scheme = "http"
	}

	// all okay
	return nil
}
