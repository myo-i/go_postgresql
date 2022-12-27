package AdvancedSQLCommands

import (
	"database/sql"
	"fmt"
	"log"
)

func challengeTimeAndToChar1() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	query := "SELECT DISTINCT(TO_CHAR(payment_date, 'MONTH')) FROM payment;"
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

func challengeTimeAndToChar2() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	// 月曜日に何回支払いがあったかを表示するクエリ
	// ↓自分の回答
	query := "SELECT COUNT(amount) FROM payment WHERE TO_CHAR(payment_date, 'D') = '2';"
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

func ChallengeTimeAndToChar() {
	challengeTimeAndToChar1()
	challengeTimeAndToChar2()
}
