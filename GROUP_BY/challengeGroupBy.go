package GROUP_BY

import (
	"database/sql"
	"fmt"
	"log"
)

type TransactionResults struct {
	StaffId int
	Sum     float64
}

type FileRatingPayment struct {
	Rating  string
	Average float64
}

type CustomerPayRank struct {
	CustomerId int
	Amount     float64
}

func challengeGroupBy1() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どのスタッフが取引を多く行ったかわかるクエリ
	rows, err := db.Query("SELECT staff_id, COUNT(amount) FROM payment GROUP BY staff_id ORDER BY SUM(amount);")
	var pp []TransactionResults
	for rows.Next() {
		var p TransactionResults
		err := rows.Scan(&p.StaffId, &p.Sum)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func challengeGroupBy2() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どの映画のレートが平均で稼いでいるかわかるクエリ
	rows, err := db.Query("SELECT rating, ROUND(AVG(replacement_cost), 2) FROM film GROUP BY rating ORDER BY AVG(replacement_cost) DESC;\n")
	var pp []FileRatingPayment
	for rows.Next() {
		var p FileRatingPayment
		err := rows.Scan(&p.Rating, &p.Average)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func challengeGroupBy3() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=dvdrental sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// どの日にちにどれくらいの決済があったかを表示するクエリ
	rows, err := db.Query("SELECT customer_id, ROUND(SUM(amount), 2) FROM payment GROUP BY customer_id ORDER BY SUM(amount) DESC LIMIT 5;")
	var pp []CustomerPayRank
	for rows.Next() {
		var p CustomerPayRank
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

func ChallengeGroupBy() {
	challengeGroupBy1()
	challengeGroupBy2()
	challengeGroupBy3()
}
