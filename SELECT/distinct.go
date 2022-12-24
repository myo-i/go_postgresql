package SELECT

import (
	"database/sql"
	"fmt"
	"log"
)

var rating []string

func ChallengeDistinct() {

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	cmd := "SELECT DISTINCT rating FROM film;"
	rows, err := db.Query(cmd)
	for rows.Next() {
		var p string
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		rating = append(rating, p)
	}
	for _, r := range rating {
		fmt.Println(r)
	}
}
