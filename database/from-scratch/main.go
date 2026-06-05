package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./library.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec(`CREATE TABLE IF NOT EXIST authors(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
		);`)
	db.Exec(`CREATE TABLE IF NOT EXISTS books(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author_id INTEGER,
		FOREIGN KEY(author_id) REFERENCES authors(id)
		);`)

	fmt.Println("Библиотека создана")
}
