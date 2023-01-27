/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"database/sql"
	"errors"

	"github.com/spf13/cobra"
	"github.com/rs/zerolog"
	"github.com/TheSnakeWitcher/mypeople/contacts"
)

var (
	Logger  zerolog.Logger
    Queries *contacts.Queries
	TermLogger = zerolog.New(os.Stdout)
)

var (
    UnknowErr = errors.New("unknow")
    NoArgErr = errors.New("not args provided")
)

var(
    jsonFlag bool
)


// checkNotArg returns true if args is empty
func checkNoArg(args []string) bool {
    if len(args) == 0 {
        return true
    }
    return false
}

// checkSingleArg returns true if args contain a single element
func checkSingleArg(args []string) bool {
    if len(args) == 1 {
        return true
    }
    return false
}

// checkMultiArg returns true if args contain multiples elements
func checkMultiArg(args []string) bool {
    if len(args) > 1 {
        return true
    }
    return false
}


var rootCmd = &cobra.Command{
	Use:   "mypeople",
	Short: "A brief description of your application",
	Long: `long doc here`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}


func Execute(db *sql.DB,logger zerolog.Logger) {
	Logger = logger
	Queries = contacts.New(db)

 
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&jsonFlag,"json","j",false,"output as json")
}
