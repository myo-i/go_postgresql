package SELECT

import (
	"database/sql"
	"fmt"
	"log"
)

func like1() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT first_name, last_name FROM customer WHERE first_name ILIKE 'j%' AND last_name LIKE 'S%';")
	var pp []Name
	for rows.Next() {
		var p Name
		err := rows.Scan(&p.FirstName, &p.LastName)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func like2() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT first_name FROM customer WHERE first_name LIKE '_her%';")
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

func like3() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT first_name, last_name FROM customer WHERE first_name LIKE 'A%' AND last_name NOT LIKE 'B%' ORDER BY last_name;")
	var pp []Name
	for rows.Next() {
		var p Name
		err := rows.Scan(&p.FirstName, &p.LastName)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func LikeILike() {
	like1()
	like2()
	like3()
}
