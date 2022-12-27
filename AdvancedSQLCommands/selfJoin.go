package AdvancedSQLCommands

import (
	"database/sql"
	"fmt"
	"log"
)

type SameFilmLength struct {
	Title1 string
	Title2 string
	Length int
}

func selfJoin() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	// 同じテーブル内で映画の上映時間が同じ物をペアとして表示
	query := "SELECT f1.title, f2.title, f1.length FROM film AS f1 INNER JOIN film AS f2 ON f1.film_id != f2.film_id AND f1.length = f2.length ORDER BY length;"
	// めっちゃ長いから改行したやつおいておく
	//SELECT f1.title, f2.title, f1.length
	//FROM film AS f1
	//INNER JOIN film AS f2 ON
	//f1.film_id != f2.film_id
	//AND f1.length = f2.length;
	rows, err := db.Query(query)
	var pp []SameFilmLength
	for rows.Next() {
		var p SameFilmLength
		err := rows.Scan(&p.Title1, &p.Title2, &p.Length)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func SelfJoin() {
	selfJoin()
}
