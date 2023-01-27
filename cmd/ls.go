/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)


var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list contacts",
	Long: `additional filter options here`,
	Run: func(cmd *cobra.Command, args []string) {
	    logger := Logger.With().Logger()
	    data ,err := Queries.ListPeople(context.Background())
	    if err != nil {
	        logger.Error().Err(errors.New("failed to list people"))
	    }
	    //fmtData := make([]*base.PeopleData,len(data))
	    //for i,_ := range data {
	    //    fmtData[i] = data[i].ToPeopleData()
	    //}

	    if jsonFlag {
	        printData , err := json.MarshalIndent(data,"","\t")
	        if err != nil {
	           logger.Error().Err(err).Msg("failed to list output as json")
	        }
	        fmt.Println(string(printData))
	        return
	    }

	    fmt.Println(data)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
