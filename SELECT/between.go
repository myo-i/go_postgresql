package SELECT

import (
	"database/sql"
	"fmt"
	"log"
)

var amounts []float64
var paymentDate []string

func betweenInt() {
	//connStr := "postgresql://postgres:password@::1/todos?sslmode=disable"

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT amount FROM payment WHERE amount BETWEEN 8 AND 9")
	//var pp []Name
	for rows.Next() {
		var p float64
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		amounts = append(amounts, p)
	}
	for _, p := range amounts {
		fmt.Println(p)
	}
}

func betweenDate() {

	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("SELECT payment_date FROM payment WHERE payment_date BETWEEN '2007-02-01' AND '2007-02-15';")
	//var pp []Name
	for rows.Next() {
		var p string
		err := rows.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		paymentDate = append(paymentDate, p)
	}
	for _, p := range paymentDate {
		fmt.Println(p)
	}
}

// Between 8 AND 9であれば8以上9以下を表すのに対し、日にちに対してBETWEEN '2007-02-01' AND '2007-02-14'と記述すると
// 2007-02-14のデータは取得されない。 なので02-16までのデータが欲しい場合、BETWEEN '2007-02-01' AND '2007-02-17'と記述する必要がある
func Between() {
	betweenInt()
	betweenDate()
}
