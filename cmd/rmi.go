package cmd

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/spf13/cobra"
)

// rmiCmd represents the rmi command
var rmiCmd = &cobra.Command{
	Use:   "rmi",
	Short: "deletes the docker images",
	Long:  `deletes the docker images`,
	Run:   removeImages,
}

func init() {
	RootCmd.AddCommand(rmiCmd)

	// Here you will define your flags and configuration settings.
	rmiCmd.Flags().BoolP("untagged", "u", true, "untagged images")
}

func removeImages(cmd *cobra.Command, args []string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal("Unable to communicate with docker daemon")
	}
	ctx := context.Background()

	filters := filters.NewArgs()
	exited, _ := cmd.Flags().GetBool("untagged")
	if exited {
		filters.Add("dangling", "true")
	}

	options := types.ImageListOptions{Filters: filters}
	imageList, err := cli.ImageList(ctx, options)
	if err != nil {
		log.Fatal(err)
	}

	dryrun, _ := cmd.Flags().GetBool("dryrun")
	for _, image := range imageList {
		fmt.Print("Removing ", image.ID[7:16], " ... ")

		if !dryrun {
			_, err := cli.ImageRemove(ctx, image.ID, types.ImageRemoveOptions{})
			if err != nil {
				fmt.Print(err)
			} else {
				fmt.Print("\tRemoved")
			}
		}
		fmt.Println("")

	}

}
