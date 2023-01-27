package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/TheSnakeWitcher/mypeople/contacts/base"
)


var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list contacts",
	Long: `additional filter options here`,
	Run: func(cmd *cobra.Command, args []string) {
	    if !checkNoArg(args) {
	        fmt.Println("error: arguments provided")
	        return
	    }

	    logger := Logger.With().Logger()
	    contacts ,err := Queries.ListPeople(context.Background())
	    if err != nil {
	        logger.Error().Err(errors.New("failed to list people"))
	    }

	    fmtData := make([]*base.PeopleData,len(contacts))
	    for i , contact := range contacts {
	        fmtData[i] = contact.ToPeopleData()
	    }

	    if jsonFlag {
	        printData , err := json.MarshalIndent(fmtData,"","\t")
	        if err != nil {
	           logger.Error().Err(err).Msg("failed to list output as json")
	        }
	        fmt.Println(string(printData))
	        return
	    }

	    fmt.Println(contacts)
	},
}


func init() {
	rootCmd.AddCommand(lsCmd)
}
