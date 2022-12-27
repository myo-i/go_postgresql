package AdvancedSQLCommands

import (
	"database/sql"
	"fmt"
	"log"
)

func math1() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	query := "SELECT ROUND(rental_rate/replacement_cost, 2) FROM film;"
	rows, err := db.Query(query)
	for rows.Next() {
		var p float64
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(p)
	}
}

func math2() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	query := "SELECT 0.1 * replacement_cost AS deposit FROM film;"
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

func Math() {
	math1()
	math2()
}
