package SELECT

import (
	"database/sql"
	"fmt"
	"log"
)

var firstName []string

func in1() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT DISTINCT amount FROM payment WHERE amount NOT IN (0.99, 1.98, 1.99) ORDER BY amount;")
	//var pp []Name
	for rows.Next() {
		var p float64
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		amounts = append(amounts, p)
	}
	for _, p := range amounts {
		fmt.Println(p)
	}
}

func in2() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT first_name FROM customer WHERE first_name IN ('John', 'Jake', 'Julie');")
	//var pp []Name
	for rows.Next() {
		var p string
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		firstName = append(firstName, p)
	}
	for _, p := range firstName {
		fmt.Println(p)
	}
}

func In() {
	in1()
	in2()
}
