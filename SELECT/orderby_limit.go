package SELECT

import (
	"database/sql"
	"fmt"
	"log"
)

type StoreCus struct {
	FirstName string
	LastName  string
}

type RecentPurchase struct {
	PaymentId  int
	CustomerId int
	StaffId    int
	RentalId   int
	Amount     float64
	// 日付型はstringで処理すると別の文字に変換される
	PaymentDate string
}

func OrderBy() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// カラムを選択していない場合でもソートできる
	rows, err := db.Query("SELECT first_name, last_name FROM customer ORDER BY store_id DESC, first_name ASC;")
	var pp []StoreCus
	for rows.Next() {
		var p StoreCus
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

func Limit() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// カラムを選択していない場合でもソートできる
	rows, err := db.Query("SELECT * FROM payment WHERE amount != 0.00 ORDER BY payment_date LIMIT 5;")
	var pp []RecentPurchase
	for rows.Next() {
		var p RecentPurchase
		err := rows.Scan(&p.PaymentId, &p.CustomerId, &p.StaffId, &p.RentalId, &p.Amount, &p.PaymentDate)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}

	for _, p := range pp {
		fmt.Println(p)
	}

}

func OrderByLimit() {
	OrderBy()
	Limit()
}
