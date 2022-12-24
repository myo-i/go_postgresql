package JOINS

import (
	"database/sql"
	"fmt"
	"log"
)

func fullOuterJoin() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// FULL OUTER JOINはベン図でいうと、一致するものから片方にしかないものまで全て取得する
	// 下記のクエリは顧客情報はあるが支払い情報がないもの、またはその逆を取得するクエリだが結果が0になる。登録されている全ての顧客は必ず支払いをしているため。
	query := "SELECT COUNT(*) FROM customer FULL OUTER JOIN payment ON customer.customer_id = payment.customer_id WHERE customer.customer_id IS null OR payment.payment_id IS null;"
	rows, err := db.Query(query)
	for rows.Next() {
		var p int
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(p)
	}
}

func FullOuterJoin() {
	fullOuterJoin()
}
