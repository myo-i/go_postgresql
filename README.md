# go_postgresql

## sql: unknown driver "postgres" (forgotten import?) が発生してposgreに接続できない
インポートした"database/sql"はあくまで汎用的なインターフェースであり、具体的な実装をインポートする必要がある。  
そのため、インポートに[ _ "github.com/lib/pq" ]（先頭の"_"は使用しなくても残すという意味）を追加するとDBに接続できた。  
参考記事  
https://stackoverflow.com/questions/52789531/how-do-i-solve-panic-sql-unknown-driver-postgres-forgotten-import
