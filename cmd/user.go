package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "handle user subcommands",
	Long: `create, update(token information)and delete user

see official document (https://docs.pixe.la) for more detail`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: change this statement print `pixela user` subcommands
		fmt.Println("user called")
	},
}

func init() {
	rootCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
