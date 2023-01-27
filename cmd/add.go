/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/TheSnakeWitcher/mypeople/contacts"
	"github.com/spf13/cobra"
)

var (
    picPath string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new contact",
	Long:  `long description here`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := Logger.With().Logger()

		switch {

		case checkNoArg(args):
			logger.Error().Err(NoArgErr)
		    fmt.Println("error: ",NoArgErr)

		case checkSingleArg(args):
			people, err := Queries.AddPeople(
			    context.Background(),
                contacts.AddPeopleParams{
                    args[0],
                    picPath,
                },
            )
			if err != nil {
				logger.Error().Err(err).Msg("add people op")
			}
			fmt.Println("people added: ", people)

		case checkMultiArg(args):
			for i, _ := range args {
				_, err := Queries.AddPeople(
				    context.Background(),
				    contacts.AddPeopleParams{args[i],picPath},
				)
				if err != nil {
					logger.Error().Err(err).Msg("add people operation ")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&picPath,"pic","p","","path to contact picture")
}
