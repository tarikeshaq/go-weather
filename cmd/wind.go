package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tarikeshaq/go-weather/api"
)

func getWindCmd() *cobra.Command {
	windCmd := &cobra.Command{
		Use:   "wind",
		Short: "retrieves the temprature in a given city",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			wind, err := api.GetWind(args[0])
			if err != nil {
				log.Fatalf("Could not read wind: %v", err)
			}
			cmd.Printf("The wind rate at %s is %v", args[0], wind)
		},
	}
	return windCmd
}
