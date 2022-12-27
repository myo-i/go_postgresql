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

func timestamp2() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	// タイムスタンプから念を抽出
	//query := "SELECT EXTRACT(YEAR FROM payment_date) AS createmonth FROM payment;"

	// タイムスタンプから現在までの年月日を計算
	query := "SELECT AGE(payment_date) FROM payment;"
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

func toChar1() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	// 年月日をフォーマットを指定して出力
	//query := "SELECT TO_CHAR(payment_date, 'MONTH-YYYY') FROM payment;"
	//query := "SELECT TO_CHAR(payment_date, 'mon/dd/YYYY') FROM payment;"
	query := "SELECT TO_CHAR(payment_date, 'MM/dd/YYYY') FROM payment;"
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
	//timestamp1()
	//timestamp2()
	toChar1()
}
