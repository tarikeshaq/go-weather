package cmd

import (
	"log"
	"strings"

	"github.com/tarikeshaq/go-weather/api"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "go-weather",
		Short: "A fast tool used to retrieve the weather",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			main, err := api.GetMain(strings.Join(args[:], "+"))
			if err != nil {
				log.Fatalf("Unable to retrieve weather: %v", err)
			}
			cmd.Printf("The weather today is: %s", main)
		},
	}
	tempCmd := getTempCmd()
	windCmd := getWindCmd()
	rootCmd.AddCommand(tempCmd, windCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Unable to execute: %v", err)
	}
}
