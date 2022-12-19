package SELECT

import (
	"database/sql"
	"fmt"
	"log"
)

var emails []string
var descriptions []string
var phone []int

func problem1(db *sql.DB) {
	cmd := "SELECT email FROM customer WHERE first_name = 'Nancy' AND last_name = 'Thomas';"
	rows, err := db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	var email string
	addSlice(rows, emails, email)
}

func problem2(db *sql.DB) {
	cmd := "SELECT description FROM film WHERE title ='Outlaw Hanky';"
	rows, err := db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	var description string
	addSlice(rows, descriptions, description)
}

func problem3(db *sql.DB) {
	cmd := "SELECT phone FROM address WHERE address = '259 Ipoh Drive';"
	rows, err := db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	var number int
	addSlice(rows, phone, number)
}

// addSlice は取得した全業のデータをスライスに格納します
func addSlice[T string | int](rows *sql.Rows, slice []T, row T) {
	for rows.Next() {
		err := rows.Scan(&row)
		if err != nil {
			log.Fatalln(err)
		}
		slice = append(slice, row)
	}
	for _, p := range slice {
		fmt.Println(p)
	}
}

func ChallengeWhere() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	problem1(db)
	problem2(db)
	problem3(db)
}
