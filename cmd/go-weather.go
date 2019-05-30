package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "go-weather",
		Short: "A fast tool used to retrieve the weather",
		Long:  ``,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}
	tempCmd := getTempCmd()
	windCmd := getWindCmd()
	rootCmd.AddCommand(tempCmd, windCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Unable to execute: %v", err)
	}
}
