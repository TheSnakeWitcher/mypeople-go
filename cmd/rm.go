package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a people",
	Long:  `remove a people`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := Logger.With().Logger()
		ctx := context.Background()

		switch {

		case checkNoArg(args):
			logger.Error().Err(NoArgErr)
		    fmt.Println("error: ",NoArgErr)

		case checkSingleArg(args):
			name := args[0]
			err := Queries.DelPeople(ctx, name)
			if err != nil {
				logger.Error().Str("name", name).Err(errors.New("failed to delete people"))
			}
			logger.Trace().Str("name", name)
			fmt.Println("deleted people: ", name)
		case checkMultiArg(args):
			for _, name := range args {
				err := Queries.DelPeople(ctx, name)
				if err != nil {
					logger.Error().Str("name", name).Err(errors.New("failed to delete people"))
				}
				logger.Trace().Str("name", name)
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
