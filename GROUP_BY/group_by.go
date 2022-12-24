package GROUP_BY

import (
	"database/sql"
	"fmt"
	"log"
)

type CustomerPurchaseLog struct {
	CustomerId int
	Sum        float64
}

type PayToStaff struct {
	CustomerId int
	StaffId    int
	Sum        float64
}

type PayPerDate struct {
	Date string
	Sum  float64
}

func groupby1() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どの顧客が最も支払った金額が多いのかわかるクエリ
	//rows, err := db.Query("SELECT customer_id, SUM(amount) FROM payment GROUP BY customer_id ORDER BY SUM(amount) DESC;")
	rows, err := db.Query("SELECT customer_id, COUNT(amount) FROM payment GROUP BY customer_id ORDER BY SUM(amount) DESC;")
	var pp []CustomerPurchaseLog
	for rows.Next() {
		var p CustomerPurchaseLog
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

func groupby2() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どの顧客がどのスタッフに対してどのくらいの金額を支払ったのかわかるクエリ
	rows, err := db.Query("SELECT customer_id, staff_id, SUM(amount) FROM payment GROUP BY customer_id, staff_id ORDER BY customer_id;")
	var pp []PayToStaff
	for rows.Next() {
		var p PayToStaff
		err := rows.Scan(&p.CustomerId, &p.StaffId, &p.Sum)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func groupbyDate() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どの日にちにどれくらいの決済があったかを表示するクエリ
	rows, err := db.Query("SELECT DATE(payment_date), SUM(amount) FROM payment GROUP BY DATE(payment_date) ORDER BY SUM(amount) DESC;")
	var pp []PayPerDate
	for rows.Next() {
		var p PayPerDate
		err := rows.Scan(&p.Date, &p.Sum)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p.Date[0:10], p.Sum)
	}
}

func GroupBy() {
	//groupby1()
	//groupby2()
	groupbyDate()
}
