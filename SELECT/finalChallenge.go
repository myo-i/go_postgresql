package SELECT

import (
	"database/sql"
	"fmt"
	"log"
)

func challenge1() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT COUNT(amount) FROM payment WHERE amount >= 5.00;")
	//var pp []Name
	for rows.Next() {
		var p int
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(p)
	}
}

func challenge2() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT COUNT(*) FROM actor WHERE first_name LIKE 'P%';")
	//var pp []Name
	for rows.Next() {
		var p string
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(p)
	}
}

func challenge3() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT COUNT(DISTINCT(district))  FROM address;")
	//var pp []Name
	for rows.Next() {
		var p string
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(p)
	}
}

func challenge4() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT COUNT(*) FROM film WHERE rating = 'R'  AND replacement_cost BETWEEN 5 AND 15;")
	//var pp []Name
	for rows.Next() {
		var p string
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(p)
	}
}

func challenge5() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT COUNT(*) FROM film WHERE title LIKE '%Truman%';")
	//var pp []Name
	for rows.Next() {
		var p string
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(p)
	}
}

func FinalChallenge() {
	challenge1()
	challenge2()
	challenge3()
	challenge4()
	challenge5()
}
