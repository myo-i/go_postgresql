package SELECT

import (
	"database/sql"
	"fmt"
	"log"
)

type Customer struct {
	FirstName string
	LastName  string
	Email     string
}

func ChallengeSelect() {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	cmd := "SELECT first_name, last_name, email FROM customer;"
	rows, err := db.Query(cmd)
	var pp []Customer
	for rows.Next() {
		var p Customer
		err := rows.Scan(&p.FirstName, &p.LastName, &p.Email)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}
