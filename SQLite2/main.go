package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID int64
	Name string
	Age  int64
}

func main() {
	// SQLiteデータベースに接続
	db, err := sql.Open("sqlite3", "./example2.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// テーブル作成
	createTableSQL := `CREATE TABLE IF NOT EXISTS user (
	    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			age INTEGER NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("テーブルが作成されました")

	// レコードの挿入
	users := []*User{
		{Name: "ten", Age: 32},
		{Name: "Gop", Age: 10},
	}

	for i := range users {
		const insertSQL = "INSERT INTO user(name, age) VALUES (?, ?)"
		result, err := db.Exec(insertSQL, users[i].Name, users[i].Age)
		if err != nil {
			log.Fatal(err)
		}

		//挿入されたレコードのIDを取得
		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		users[i].ID = id
		fmt.Printf("レコードが挿入されました: %v\n", users[i])
	}

	// 複数レコードのスキャン
	rows, err := db.Query("SELECT id, name, age FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("ユーザーリスト:")
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", u.ID, u.Name, u.Age)
	}

	// スキャン中にエラーが発生していないか確認
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// レコードの更新
	const updateSQL = "UPDATE user SET age= age + 1 WHERE id = ?"
	result, err := db.Exec(updateSQL, 1) // IDが1のユーザーの年齢を更新
	if err != nil {
		log.Fatal(err)
	}

	// 更新されたレコード数を取得
	rowsupdated, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("更新されたレコード数: %d\n", rowsupdated)

	// 最後に再度データを表示
	rows, err = db.Query("SELECT id, name, age FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("更新後のユーザーリスト:")
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", u.ID, u.Name, u.Age)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}