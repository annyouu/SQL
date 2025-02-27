package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func main() {
	// SQLiteデータベースを作成
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// テーブル作成クエリ
	const createTableSQL = `CREATE TABLE IF NOT EXISTS users (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		  name TEXT,
			age INTEGER
	);`

	// テーブルの作成
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// テーブル作成完了のメッセージ
	fmt.Println("テーブルusersが作成されました。")

	// データを挿入する
	insertUserSQL := `INSERT INTO users (name, age) VALUES (?,?)`
	_, err = db.Exec(insertUserSQL, "tarou", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("データが挿入されました。")
	
	// テーブルからデータを読み取る
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 結果を表示
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name:%s, Age:%d\n", id, name, age)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}


// func main() {
// 	// SQLiteデータベースを作成
// 	db, err := sql.Open("sqlite3", "test.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// テーブル作成
// 	sqlStmt := `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)`;
// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		log.Fatalf("テーブル作成エラー: %s\n", err)
// 	}

// 	// データを追加
// 	_, err = db.Exec("INSERT INTO users (name) VALUES (?)", "tarou")
// 	if err != nil {
// 		log.Fatalf("データ挿入エラー: %s\n", err)
// 	}
	
// 	// データを取得
// 	rows, err := db.Query("SELECT id, name FROM users")
// 	if err != nil {
// 		log.Fatalf("データ取得エラー: %s\n", err)
// 	}
// 	defer rows.Close()

// 	fmt.Println("データ一覧:")
// 	for rows.Next() {
// 		var id int
// 		var name string
// 		err = rows.Scan(&id, &name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("ID: %d, Name: %s\n", id, name)
// 	}
// }


