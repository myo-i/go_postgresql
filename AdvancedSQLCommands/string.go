package AdvancedSQLCommands

import (
	"database/sql"
	"fmt"
	"log"
)

func string1() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	// 文字列の連結
	//query := "SELECT first_name || ' ' || last_name AS full_name FROM customer;"
	query := "SELECT upper(first_name) || ' ' || upper(last_name) AS full_name FROM customer;"
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

func string2() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	query := "SELECT LOWER(LEFT(first_name, 1)) || LOWER(last_name) || '@gmail.com'  AS custom_email FROM customer;"
	//query := "SELECT COUNT(*) FROM payment WHERE EXTRACT(dow FROM payment_date) = 1;"
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

func String() {
	//string1()
	string2()
}
