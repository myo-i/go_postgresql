package GROUP_BY

import (
	"database/sql"
	"fmt"
	"log"
)

type CustomerPayLog struct {
	CustomerId  int
	AboutAmount float64
}

func challengeHaving1() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どの顧客がどのスタッフに対してどのくらいの金額を支払ったのかわかるクエリ
	rows, err := db.Query("SELECT customer_id, COUNT(amount) FROM payment GROUP BY customer_id HAVING COUNT(amount) >= 40;")
	var pp []CustomerPayLog
	for rows.Next() {
		var p CustomerPayLog
		err := rows.Scan(&p.CustomerId, &p.AboutAmount)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func challengeHaving2() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どの顧客がどのスタッフに対してどのくらいの金額を支払ったのかわかるクエリ
	rows, err := db.Query("SELECT customer_id, SUM(amount) FROM payment WHERE staff_id = 2 GROUP BY customer_id HAVING SUM(amount) >= 100;")
	var pp []CustomerPayLog
	for rows.Next() {
		var p CustomerPayLog
		err := rows.Scan(&p.CustomerId, &p.AboutAmount)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func ChallengeHaving() {
	challengeHaving1()
	challengeHaving2()
}
