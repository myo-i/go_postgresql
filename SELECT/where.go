package SELECT

import (
	"database/sql"
	"fmt"
	"log"
)

var email []string
var description []string
var phone []int

func problem1() {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	cmd := "SELECT email FROM customer WHERE first_name = 'Nancy' AND last_name = 'Thomas';"
	rows, err := db.Query(cmd)
	for rows.Next() {
		var p string
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		email = append(email, p)
	}
	for _, p := range email {
		fmt.Println(p)
	}
}

func problem2() {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	cmd := "SELECT description FROM film WHERE title ='Outlaw Hanky';"
	rows, err := db.Query(cmd)
	for rows.Next() {
		var p string
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		description = append(description, p)
	}
	for _, d := range description {
		fmt.Println(d)
	}
}

func problem3() {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	cmd := "SELECT phone FROM address WHERE address = '259 Ipoh Drive';"
	rows, err := db.Query(cmd)
	for rows.Next() {
		var p int
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		phone = append(phone, p)
	}
	for _, p := range phone {
		fmt.Println(p)
	}
}

func ChallengeWhere() {
	problem1()
	problem2()
	problem3()
}
