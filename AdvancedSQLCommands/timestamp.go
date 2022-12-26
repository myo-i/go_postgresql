package AdvancedSQLCommands

import (
	"database/sql"
	"fmt"
	"log"
)

func timestamp1() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	//query := "SHOW TIMEZONE;"
	//query := "SELECT NOW();"
	//query := "SELECT TIMEOFDAY();"
	//query := "SELECT CURRENT_DATE;"
	query := "SELECT CURRENT_TIME;"
	rows, err := db.Query(query)
	for rows.Next() {
		var p string
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(p)
	}
}

func TimeStamp() {
	timestamp1()
}
