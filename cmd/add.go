/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var ( picPath string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new people",
	Long:  `long description here`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := Logger.With().Logger()

		switch {

		case checkNoArg(args):
			logger.Error().Err(NoArgErr)
		    fmt.Println("error: ",NoArgErr)

		case checkSingleArg(args):
			people, err := Queries.AddPeople(context.Background(), args[0])
			if err != nil {
				logger.Error().Err(err).Msg("add people op")
			}
			fmt.Println("people added: ", people)

		case checkMultiArg(args):
			for i, _ := range args {
				_, err := Queries.AddPeople(context.Background(), args[i])
				if err != nil {
					logger.Error().Err(err).Msg("add people operation ")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

    // set args
	//addCmd.SetArgs([]string{"pic"})

    // persist to all subcommands
	// addCmd.PersistentFlags().String("foo", "", "A help for foo") 

    // flags which will only run when this command
	//addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle") 
	addCmd.Flags().StringVarP(&picPath,"pic","p","","picture")

}
