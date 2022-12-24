package GROUP_BY

import (
	"database/sql"
	"fmt"
	"log"
)

type CustomerPayment struct {
	CustomerId int
	Amount     float64
}

type StorePayment struct {
	StoreId int
	Amount  float64
}

func having1() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どの顧客がどのスタッフに対してどのくらいの金額を支払ったのかわかるクエリ
	rows, err := db.Query("SELECT customer_id, SUM(amount) FROM payment GROUP BY customer_id HAVING SUM(amount) > 100;")
	var pp []CustomerPayment
	for rows.Next() {
		var p CustomerPayment
		err := rows.Scan(&p.CustomerId, &p.Amount)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func having2() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どの顧客がどのスタッフに対してどのくらいの金額を支払ったのかわかるクエリ
	rows, err := db.Query("SELECT store_id, COUNT(*) FROM customer GROUP BY store_id HAVING COUNT(*) > 300;")
	var pp []StorePayment
	for rows.Next() {
		var p StorePayment
		err := rows.Scan(&p.StoreId, &p.Amount)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func Having() {
	having1()
	having2()
}
