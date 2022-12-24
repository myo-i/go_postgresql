package JOINS

import (
	"database/sql"
	"fmt"
	"log"
)

type TellChangeTaxLaws struct {
	District string
	Email    string
}

type ActorsWork struct {
	Title     string
	FirstName string
	LastName  string
}

func challengeJoins1() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln("SQL Open Error has occured", err)
	}

	// FULL OUTER JOINはベン図でいうと、一致するものから片方にしかないものまで全て取得する
	// 下記のクエリは顧客情報はあるが支払い情報がないもの、またはその逆を取得するクエリだが結果が0になる。登録されている全ての顧客は必ず支払いをしているため。
	query := "SELECT district, email FROM address INNER JOIN customer ON address.address_id = customer.address_id WHERE district = 'California';"
	rows, err := db.Query(query)
	var pp []TellChangeTaxLaws
	for rows.Next() {
		var p TellChangeTaxLaws
		err := rows.Scan(&p.District, &p.Email)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

// 応用的な問題だったので要復習だけど全然できた！！
func challengeJoins2() {
	db, err := sql.Open("postgres", DataSource)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// 完全初見でできたけど頭の中で何を考えていたのかwを書いてみる
	// まず、最終的にNick Wahlbergが出演した作品のタイトルが欲しかった
	// つまりアウトプットとしてはtitle＋おまけにfirst_name, last_nameが欲しい
	// これでテーブルとしてはfilmとactorが必要なのがわかる
	// しかし、filmには誰が出演した作品なのかが記載されておらず、代わりにfilm_actorテーブルにactor_idとfilm_idが存在していた。
	// 以上のことからactorテーブルとfilm_actorテーブルを使って、Nick Wahlbergのactor_idを持つfilm_idを割り出せれば、film_idを使ってfilmからtitleを取得できると考えた。
	// それを実現するために「actor.actor_id = film_actor.actor_idが成り立つテーブルを作成して、そのテーブルにfilm.film_id = film_actor.film_idを成り立たせて、
	//　WHERE first_name = 'Nick' AND last_name = 'Wahlberg'を条件を絞れば行けんじゃね？」と考えてクエリを作成した

	// 最初は下記のクエリを作成したがちょっとわかりにくかったから変えた
	//query := "SELECT district, email FROM address INNER JOIN customer ON address.address_id = customer.address_id WHERE district = 'California';"

	query := "SELECT title, first_name, last_name FROM actor INNER JOIN film_actor ON actor.actor_id = film_actor.actor_id INNER JOIN film ON film.film_id = film_actor.film_id WHERE first_name = 'Nick' AND last_name = 'Wahlberg';"
	rows, err := db.Query(query)
	var pp []ActorsWork
	for rows.Next() {
		var p ActorsWork
		err := rows.Scan(&p.Title, &p.FirstName, &p.LastName)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
}

func ChallengeJoins() {
	challengeJoins1()
	challengeJoins2()
}
