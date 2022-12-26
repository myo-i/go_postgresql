# go_postgresql

## sql: unknown driver "postgres" (forgotten import?) が発生してposgreに接続できない
インポートした"database/sql"はあくまで汎用的なインターフェースであり、具体的な実装をインポートする必要がある。  
そのため、インポートに[ _ "github.com/lib/pq" ]（先頭の"_"は使用しなくても残すという意味）を追加するとDBに接続できた。  
参考記事  
https://stackoverflow.com/questions/52789531/how-do-i-solve-panic-sql-unknown-driver-postgres-forgotten-import

## DBのポート番号やホスト、ポート番号などの確認方法
参考記事  
https://stackoverflow.com/questions/58573628/how-do-you-get-the-url-of-a-postgresql-database-you-made-in-pgadmin

## GoのGenerics
参考記事  (汎用的)  
https://www.wakuwakubank.com/posts/868-go-generics/  
引数に([]int, int)や([]string, string)を使いたいときに参考になった記事  
https://eli.thegreenplace.net/2021/generic-functions-on-slices-with-go-type-parameters/  

## PostgeSQL正規表現
https://www.postgresql.org/docs/current/functions-matching.html

## postgresqlで扱えるデータ関連の関数など
https://www.postgresql.org/docs/current/functions-formatting.html
