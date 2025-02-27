package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// MySQLデータベースに接続
	dsn := "root:pass@tcp(127.0.0.1:3306)/example"
	db , err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 接続確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MySQLに接続されました")

	// テーブル作成SQL
	createTableSQL := `CREATE TABLE IF NOT EXISTS person (
	   id INT AUTO_INCREMENT PRIMARY KEY,
		 name VARCHAR(50),
		 age INT
	);`

	// テーブル作成
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("テーブルpersonが作成されました")

	// データ挿入SQL
	insertUserSQL := `INSERT INTO person (name, age) VALUES (?, ?)`
	_, err = db.Exec(insertUserSQL, "tarou", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("データが挿入されました")

	// データを読み取る
	rows, err := db.Query("SELECT id, name, age FROM person")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//結果表示
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
	// エラーチェック
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}