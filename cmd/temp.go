package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tarikeshaq/go-weather/api"
)

func getTempCmd() *cobra.Command {
	tempCmd := &cobra.Command{
		Use:   "temp",
		Short: "retrieves the tempreture at the city",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tempreture, err := api.GetTemp(strings.Join(args[:], "+"))
			if err != nil {
				log.Fatalf("Could not read tempreture: %v", err)
			}
			res := convertToC(tempreture)
			cmd.Printf("The tempreture in %s is %v", strings.Join(args[:], " "), res)
		},
	}
	return tempCmd
}

func convertToC(temp string) string {
	val, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		log.Fatalf("Could not parse float: %v", err)
	}
	val = val - 273.15

	return strconv.FormatFloat(val, 'f', 2, 64)
}
