package SELECT

import (
	"database/sql"
	"fmt"
	"log"
)

var amount []int

func Count() {

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	cmd := "SELECT COUNT(DISTINCT amount) FROM payment;"
	rows, err := db.Query(cmd)
	for rows.Next() {
		var v int
		err := rows.Scan(&v)
		if err != nil {
			log.Fatalln(err)
		}
		amount = append(amount, v)
	}
	for _, r := range amount {
		fmt.Println(r)
	}
}
