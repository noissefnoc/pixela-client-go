package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// variable for configuration file name
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pixela",
	Short: "Simple API Client for pixela (https://pixe.la/)",
	Long: `Simple API Client for pixela (https://pixe.la/).
This command can handle user, graph, pixel and webhook via API.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		rootCmd.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pixela.yaml)")
	rootCmd.PersistentFlags().StringP("username", "u", "", "pixe.la username")
	rootCmd.PersistentFlags().StringP("token", "t", "", "pixe.la user token")
	rootCmd.PersistentFlags().BoolP("verbose", "n", false, "verbose mode")

	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
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
			rootCmd.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".pixela" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".pixela")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		rootCmd.Println(err)
		os.Exit(1)
	}

	if viper.GetBool("verbose") {
		rootCmd.Println("Using config file:", viper.ConfigFileUsed())
	}
}
