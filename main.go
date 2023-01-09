package main

import (
	_ "github.com/mattn/go-sqlite3"

	"github.com/TheSnakeWitcher/mypeople/cmd"
)


var InitErr error

func InitErrCheck(err error) {
	if err != nil {
		panic(err)
	}
}


func init() { 
    InitErr = InitConfig()
	InitErrCheck(InitErr)

	InitErr = InitLogger()
	InitErrCheck(InitErr)

	InitErr = initDB(Config.DbDriver,Config.DbSource)
	InitErrCheck(InitErr)
}


func main() {
	Logger.Trace().Msg("execution starts")

    cmd.Execute(db,Logger)

	Logger.Trace().Msg("execution succed")
	LogFile.Close()
}
