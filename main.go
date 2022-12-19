package main

import (
	_ "github.com/lib/pq"
	"go_postgresql/SELECT"
)

func selectStatement() {
	SELECT.Select()
}

// main関数を実行するには実行構成の編集でディレクトリを選択し、goファイルのパッケージ名をmainにし、main関数を作成しないといけない
func main() {
	selectStatement()
}
