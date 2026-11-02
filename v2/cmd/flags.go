package cmd

import (
	"flag"
	"strings"
)

// stringSliceFlag implements flag.Value to accept multiple string values or comma-separated values.
type stringSliceFlag []string

func (s *stringSliceFlag) String() string {
	return strings.Join(*s, ",")
}

func (s *stringSliceFlag) Set(val string) error {
	for _, item := range strings.Split(val, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			*s = append(*s, item)
		}
	}
	return nil
}

// addGlobalFlags registers global flags to a FlagSet.
func addGlobalFlags(fs *flag.FlagSet, cfg *Config) {
	fs.StringVar(&cfg.ConfigFile, "config", cfg.ConfigFile, "config file (default is $HOME/.pixela.yaml)")
	fs.StringVar(&cfg.ConfigFile, "c", cfg.ConfigFile, "config file (shorthand)")

	fs.StringVar(&cfg.Username, "username", cfg.Username, "pixe.la username")
	fs.StringVar(&cfg.Username, "u", cfg.Username, "pixe.la username (shorthand)")

	fs.StringVar(&cfg.Token, "token", cfg.Token, "pixe.la user token")
	fs.StringVar(&cfg.Token, "t", cfg.Token, "pixe.la user token (shorthand)")

	fs.BoolVar(&cfg.Verbose, "verbose", cfg.Verbose, "verbose mode")
	fs.BoolVar(&cfg.Verbose, "n", cfg.Verbose, "verbose mode (shorthand)")
}

// parseFlagSet parses flags in args even if positional arguments appear before flags.
func parseFlagSet(fs *flag.FlagSet, args []string) ([]string, error) {
	var positional []string
	for len(args) > 0 {
		if err := fs.Parse(args); err != nil {
			return nil, err
		}
		rem := fs.Args()
		if len(rem) > 0 {
			positional = append(positional, rem[0])
			args = rem[1:]
		} else {
			break
		}
	}
	return positional, nil
}
