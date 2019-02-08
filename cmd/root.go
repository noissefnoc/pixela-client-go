package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"os"
	"runtime"
)

// variable for configuration file name
var (
	cui     = rwi.New() // CUI instance
	cfgFile string
)

func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	cui = ui

	rootCmd := &cobra.Command{
		Use:   "pixela",
		Short: "Simple API Client for pixela (https://pixe.la/)",
		Long: `Simple API Client for pixela (https://pixe.la/).
This command can handle user, graph, pixel and webhook via API.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	// global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pixela.yaml)")
	rootCmd.PersistentFlags().StringP("username", "u", "", "pixe.la username")
	rootCmd.PersistentFlags().StringP("token", "t", "", "pixe.la user token")
	rootCmd.PersistentFlags().BoolP("verbose", "n", false, "verbose mode")

	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())

	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(newPixelCmd())
	rootCmd.AddCommand(newGraphCmd())
	rootCmd.AddCommand(newUserCmd())
	rootCmd.AddCommand(newWebhookCmd())

	return rootCmd
}

func Execute(cui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		// panic handling
		if r := recover(); r != nil {
			cui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)

				if !ok {
					break
				}
				cui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	// execution
	exit = exitcode.Normal

	if err := newRootCmd(cui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}

	return
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			cui.OutputErrln(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".pixela" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".pixela")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		cui.OutputErrln(err)
		os.Exit(1)
	}

	if viper.GetBool("verbose") {
		cui.Outputln(fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
	}
}
