package assessmentTest

import (
	"database/sql"
	"fmt"
	"log"
)

type CustomerPayResult struct {
	CustomerId int
	Sum        float64
}

type Name struct {
	FirstName string
	LastName  string
}

func test1() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どの顧客がどのスタッフに対してどのくらいの金額を支払ったのかわかるクエリ
	rows, err := db.Query("SELECT customer_id, SUM(amount) FROM payment WHERE staff_id = 2 GROUP BY customer_id HAVING SUM(amount) >= 110;")
	var pp []CustomerPayResult
	for rows.Next() {
		var p CustomerPayResult
		err := rows.Scan(&p.CustomerId, &p.Sum)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func test2() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どの顧客がどのスタッフに対してどのくらいの金額を支払ったのかわかるクエリ
	rows, err := db.Query("SELECT COUNT(*) FROM film WHERE title LIKE 'J%';")
	for rows.Next() {
		var p int
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(p)
	}
}

func test3() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どの顧客がどのスタッフに対してどのくらいの金額を支払ったのかわかるクエリ
	rows, err := db.Query("SELECT first_name, last_name FROM customer WHERE first_name LIKE 'E%' AND address_id < 500 ORDER BY customer_id DESC LIMIT 1;")
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

func AssessmentTest1() {
	test1()
	test2()
	test3()
}
