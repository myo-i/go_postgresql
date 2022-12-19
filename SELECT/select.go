package SELECT

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Name struct {
	FirstName string
	LastName  string
}

func Select() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT first_name, last_name FROM actor")
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
