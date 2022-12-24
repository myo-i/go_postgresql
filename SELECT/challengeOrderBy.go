package SELECT

import (
	"database/sql"
	"fmt"
	"log"
)

var customerId []int

type ShortMovie struct {
	Title  string
	Length int
}

func problemOrderBy1() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// カラムを選択していない場合でもソートできる
	rows, err := db.Query("SELECT customer_id FROM payment ORDER BY payment_date ASC LIMIT 10;")
	for rows.Next() {
		var p int
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		customerId = append(customerId, p)
	}

	for _, p := range customerId {
		fmt.Println(p)
	}

}

func problemOrderBy2() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// カラムを選択していない場合でもソートできる
	rows, err := db.Query("SELECT title, length FROM film ORDER BY length ASC LIMIT 5;")
	var shortMovies []ShortMovie
	for rows.Next() {
		var p ShortMovie
		err := rows.Scan(&p.Title, &p.Length)
		if err != nil {
			log.Fatalln(err)
		}
		shortMovies = append(shortMovies, p)
	}

	for _, p := range shortMovies {
		fmt.Println(p)
	}

}

func problemOrderBy3() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// カラムを選択していない場合でもソートできる
	rows, err := db.Query("SELECT title, length FROM film WHERE length <= 50 ORDER BY length ASC;")
	var shortMovies []ShortMovie
	for rows.Next() {
		var p ShortMovie
		err := rows.Scan(&p.Title, &p.Length)
		if err != nil {
			log.Fatalln(err)
		}
		shortMovies = append(shortMovies, p)
	}

	for _, p := range shortMovies {
		fmt.Println(p)
	}

}

// 後でコード綺麗にする
func ChallengeOrderBy() {
	//problemOrderBy1()
	//problemOrderBy2()
	problemOrderBy3()
}
