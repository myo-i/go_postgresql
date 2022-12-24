package GROUP_BY

import (
	"database/sql"
	"fmt"
	"log"
)

type MinMax struct {
	Min float64
	Max float64
}

func aggMinMAX() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT MIN(replacement_cost), MAX(replacement_cost) FROM film;")
	//rows, err := db.Query("SELECT MIN(replacement_cost) FROM film;")
	//rows, err := db.Query("SELECT MAX(replacement_cost) FROM film;")
	//var pp []Name
	for rows.Next() {
		var p MinMax
		err := rows.Scan(&p.Min, &p.Max)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("minimum:%v, maximum:%v\n", p.Min, p.Max)
	}
}

func aggAverageRound() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT ROUND(AVG(replacement_cost), 2) FROM film;")
	//var pp []Name
	for rows.Next() {
		var p float64
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(p)
	}
}

func sum() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT SUM(replacement_cost) FROM film;")
	//var pp []Name
	for rows.Next() {
		var p float64
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(p)
	}
}

func Aggregation() {
	aggMinMAX()
	aggAverageRound()
	sum()
}
