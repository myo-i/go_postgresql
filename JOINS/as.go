package JOINS

import (
	"database/sql"
	"fmt"
	"log"
)

type CustomerPayment struct {
	CustomerId int
	Payment    float64
}

func as1() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// ASは変更したいカラムの直後に記述し、さらにASはどこに記述してもクエリの最後に実行される
	rows, err := db.Query("SELECT customer_id, SUM(amount) AS total_spent FROM payment GROUP BY customer_id HAVING SUM(amount) > 100;")
	var pp []CustomerPayment
	for rows.Next() {
		var p CustomerPayment
		err := rows.Scan(&p.CustomerId, &p.Payment)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func as2() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どの顧客がどのスタッフに対してどのくらいの金額を支払ったのかわかるクエリ
	rows, err := db.Query("SELECT customer_id, amount AS new_name FROM payment WHERE amount > 2;")
	var pp []CustomerPayment
	for rows.Next() {
		var p CustomerPayment
		err := rows.Scan(&p.CustomerId, &p.Payment)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func As() {
	//as1()
	as2()
}
