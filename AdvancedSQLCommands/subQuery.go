package AdvancedSQLCommands

import (
	"database/sql"
	"fmt"
	"log"
)

type OverAvgRentalRate struct {
	Name       string
	RentalRate float64
}
type Title struct {
	FilmId int
	Title  string
}
type Name struct {
	FirstName string
	LastName  string
}

func subQuery1() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	// rental_rateが平均以上のタイトルを抽出
	query := "SELECT title, rental_rate FROM film WHERE rental_rate > (SELECT AVG(rental_rate) FROM film);"
	rows, err := db.Query(query)
	var pp []OverAvgRentalRate
	for rows.Next() {
		var p OverAvgRentalRate
		err := rows.Scan(&p.Name, &p.RentalRate)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func subQuery2() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	// 返却日が2005年5月30日の作品のタイトルを取得
	// 備蓄にあったものがレンタルされた場合、inventory.inventory_id = rental.inventory_idが成り立つ
	query := "SELECT film_id, title FROM film WHERE film_id IN (SELECT inventory.film_id FROM rental INNER JOIN inventory ON inventory.inventory_id = rental.inventory_id WHERE return_date BETWEEN '2005-05-29' AND '2005-05-30') ORDER BY title;"
	// めっちゃ長いから改行したやつおいておく
	//"SELECT film_id, title
	//FROM film
	//WHERE film_id IN
	//(SELECT inventory.film_id
	//FROM rental
	//INNER JOIN inventory ON inventory.inventory_id = rental.inventory_id
	//WHERE return_date BETWEEN '2005-05-29' AND '2005-05-30')
	//ORDER BY title;"
	rows, err := db.Query(query)
	var pp []Title
	for rows.Next() {
		var p Title
		err := rows.Scan(&p.FilmId, &p.Title)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func subQuery3() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	// 支払いを11ドル以上行った人の姓と名を取得
	query := "SELECT first_name, last_name FROM customer AS c WHERE EXISTS (SELECT * FROM payment AS p WHERE p.customer_id = c.customer_id AND amount > 11);"
	// めっちゃ長いから改行したやつおいておく
	//SELECT first_name, last_name
	//FROM customer AS c
	//WHERE EXISTS
	//(SELECT * FROM payment AS p
	//WHERE p.customer_id = c.customer_id
	//AND amount > 11);
	rows, err := db.Query(query)
	var pp []Name
	for rows.Next() {
		var p Name
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

// SubQuery サブクエリは処理の順序としては、サブクエリが実行され、
// その結果に対して文の最初にあるクエリ(サブクエリのカッコ外のクエリ)が実行されるという順番
func SubQuery() {
	//subQuery1()
	//subQuery2()
	subQuery3()
}
