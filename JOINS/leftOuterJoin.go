package JOINS

import (
	"database/sql"
	"fmt"
	"log"
)

type CheckInventory struct {
	FilmId int
	Title  string
	// 今回はNULLが入ってくるとわかったているのでNULL許容で定義
	InventoryId sql.NullInt64
	StoreId     sql.NullInt64
}

func leftOuterJoin() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// FULL OUTER JOINはベン図でいうと、一致するものから片方にしかないものまで全て取得する
	// 下記のクエリは顧客情報はあるが支払い情報がないもの、またはその逆を取得するクエリだが結果が0になる。登録されている全ての顧客は必ず支払いをしているため。
	query := "SELECT film.film_id, title, inventory_id, store_id\nFROM film LEFT JOIN inventory ON inventory.film_id = film.film_id WHERE inventory.film_id IS null;"
	rows, err := db.Query(query)
	var pp []CheckInventory
	for rows.Next() {
		var p CheckInventory
		err := rows.Scan(&p.FilmId, &p.Title, &p.InventoryId, &p.StoreId)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func LeftOuterJoin() {
	leftOuterJoin()
}
