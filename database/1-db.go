package main

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main(){
	db, err := sql.Open("sqlite3", "./db")
	if err != nil{
		return fmt.Errorf("failed to open DB: %w", err)
	}
	defer db.Close()

}
