package JOINS

import (
	"database/sql"
	"fmt"
	"log"
)

func fullOuterJoin() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// FULL OUTER JOINはベン図でいうと、一致するものから片方にしかないものまで全て取得する
	// 下記のクエリは顧客情報はあるが支払い情報がないもの、またはその逆を取得するクエリだが結果が0になる。登録されている全ての顧客は必ず支払いをしているため。
	query := "SELECT * FROM customer FULL OUTER JOIN payment ON customer.customer_id = payment.customer_id WHERE customer.customer_id IS null OR payment.payment_id IS null;"
	rows, err := db.Query(query)
	var pp []PaymentCustomer
	for rows.Next() {
		var p PaymentCustomer
		err := rows.Scan(&p.PaymentId, &p.CustomerId, &p.FirstName)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func FullOuterJoin() {
	fullOuterJoin()
}
