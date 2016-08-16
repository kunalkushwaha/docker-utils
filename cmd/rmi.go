package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmiCmd represents the rmi command
var rmiCmd = &cobra.Command{
	Use:   "rmi",
	Short: "deletes the docker images",
	Long:  `deletes the docker images`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("rmi called")
	},
}

func init() {
	RootCmd.AddCommand(rmiCmd)

	// Here you will define your flags and configuration settings.
	rmiCmd.Flags().BoolP("untagged", "u", true, "untagged images")
}
