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

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "deletes docker containers in bulk",
	Long:  `deletes docker containers in bulk`,
	Run:   removeContainers,
}

func init() {
	RootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.
	rmCmd.Flags().BoolP("exited", "e", true, "exited containers")
}

func removeContainers(cmd *cobra.Command, args []string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal("Unable to communicate with docker daemon")
	}
	ctx := context.Background()

	filters := filters.NewArgs()
	exited, _ := cmd.Flags().GetBool("exited")
	if exited {
		filters.Add("status", "exited")
	}

	options := types.ContainerListOptions{Filter: filters}
	clist, err := cli.ContainerList(ctx, options)
	if err != nil {
		log.Fatal("Error : ", err)
	}

	removeOptions := types.ContainerRemoveOptions{RemoveVolumes: true, Force: true}
	dryrun, _ := cmd.Flags().GetBool("dryrun")

	for _, container := range clist {
		fmt.Print("Removing ", container.ID[0:8], " ", container.Names[0], " ... ")
		if !dryrun {
			err := cli.ContainerRemove(ctx, container.ID, removeOptions)
			if err != nil {
				fmt.Print(err)
			} else {
				fmt.Print("\tRemoved")
			}
		}
		fmt.Println("")
	}

}
