package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

	type Contact struct {
		ID          int64
		Name        string
		PhoneNumber string
	}

	// データベースに保存されている全ての連絡先を表示
	func printAllContacts(db *sql.DB) {
		rows, err := db.Query("SELECT * FROM contact")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		fmt.Println("\n現在登録されている連絡先:")
		for rows.Next() {
			var contact Contact
			err := rows.Scan(&contact.ID, &contact.Name, &contact.PhoneNumber)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("ID: %d, 名前: %s, 電話番号: %s\n", contact.ID, contact.Name, contact.PhoneNumber)

			//スキャン中にエラーが発生したらそれを表示
			if err := rows.Err(); err != nil {
				log.Fatal(err)
			}
		}
	}

	// IDを指定して連絡先の情報を更新
	func updateContact(db *sql.DB) {
		var id int
		var newName, newPhoneNumber string
		fmt.Print("更新する連絡先のIDを入力してください:")
		fmt.Scan(&id)
		fmt.Print("新しい名前を入力してください:")
		fmt.Scan(&newName)
		fmt.Print("新しい電話番号を入力してください:")
		fmt.Scan(&newPhoneNumber)

		updateSQL := "UPDATE contact SET name = ?, phone_number = ? WHERE id = ?"
		_, err := db.Exec(updateSQL, newName, newPhoneNumber, id)
		if err != nil {
			log.Fatal(err)
		}
		printAllContacts(db)
	}

	func main() {
		// SQLiteデータベースに接続
		db, err := sql.Open("sqlite3", "./phonebook.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// テーブル作成
		createTableSQL := 
		`CREATE TABLE IF NOT EXISTS contact (
		   id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			 name TEXT NOT NULL,
			 phone_number TEXT NOT NULL
		);`

		_, err = db.Exec(createTableSQL)
		if err != nil {
			log.Fatal(err)
		}

		// 起動時に現在登録されている情報を表示
		printAllContacts(db)

		// 入力モード
		for {
			var name, phoneNumber string
			fmt.Println("\n新しい連絡を追加します。名前と電話番号を入力してください。")
			fmt.Print("名前: ")
			fmt.Scan(&name)
			fmt.Print("電話番号: ")
			fmt.Scan(&phoneNumber)

			// データベースに保存
			insertSQL := "INSERT INTO contact(name, phone_number) VALUES (?, ?)"
			_, err := db.Exec(insertSQL, name, phoneNumber)
			if err != nil {
				log.Fatal(err)
			}

			// データベースの内容を表示
			printAllContacts(db)

			// 更新したいか確認
			fmt.Println("\nデータを更新しますか？ (y/n)")
			var updateChoice string
			fmt.Scan(&updateChoice)

			if updateChoice == "y" {
				updateContact(db)
			}

			// 続けて入力するか確認
			fmt.Println("\n続けて入力しますか？ (y/n)")
			var continueChoice string
			fmt.Scan(&continueChoice)
			if continueChoice != "y" {
				break
			}
		}
	}
