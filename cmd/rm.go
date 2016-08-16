package cmd

import (
	"log"

	"github.com/docker/engine-api/client"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "deletes docker containers",
	Long:  `deletes docker containers`,
	Run:   removeImages,
}

func init() {
	RootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.
	rmCmd.Flags().BoolP("exited", "e", true, "exited containers")
}

func removeImages(cmd *cobra.Command, args []string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal("Unable to communicate with docker daemon")
	}

	log.Print("API Version : ", cli.ClientVersion())

}
