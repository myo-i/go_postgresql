package JOINS

import (
	"database/sql"
	"fmt"
	"log"
)

type PaymentCustomer struct {
	PaymentId  int
	CustomerId int
	FirstName  string
}

func innerJoin1() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// INNER JOINは指定したカラムの値がテーブルA、テーブルBどちらも同じものを結合して取得できる
	rows, err := db.Query("SELECT payment_id, payment.customer_id, first_name FROM payment INNER JOIN customer ON payment.customer_id = customer.customer_id;")
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

func InnerJoin() {
	innerJoin1()
}
