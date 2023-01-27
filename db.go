package main

import "database/sql"

var db *sql.DB

func initDB(driver,source string) error {
    var err error

	db, err = sql.Open(Config.DbDriver,Config.DbSource)
	if err != nil {
		panic(err)
	}

    err = db.Ping()
	if err != nil {
		panic(err)
	}

    return nil
}
