package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/argon2"
)

type User struct {
	ID           int
	Name         string
	PasswordHash string
}

func hash(pass string) string {
	salt := make([]byte, 8)
	_, _ = rand.Read(salt)
	h := argon2.IDKey([]byte(pass), salt, 1, 64*1024, 4, 32)
	return base64.RawStdEncoding.EncodeToString(salt) + "!" + base64.RawStdEncoding.EncodeToString(h)
}

func verify(pass, stored string) bool {
	parts := strings.Split(stored, "!")
	if len(parts) < 2 {
		return false
	}
	salt, _ := base64.RawStdEncoding.DecodeString(parts[0])
	oldHash, _ := base64.RawStdEncoding.DecodeString(parts[1])
	newHash := argon2.IDKey([]byte(pass), salt, 1, 64*1024, 4, 32)
	return string(newHash) == string(oldHash)
}

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test_db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	var choice int
	var user, pass string

	fmt.Print("1. Register\n2. Login\nChoice: ")
	fmt.Scan(&choice)
	fmt.Print("Name: ")
	fmt.Scan(&user)
	fmt.Print("Password: ")
	fmt.Scan(&pass)

	if choice == 1 {
		_, err := db.Exec("INSERT INTO users (name, password_hash) VALUES (?, ?)", user, hash(pass))
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Success!")
		}
	} else {
		var stored string
		err := db.QueryRow("SELECT password_hash FROM users WHERE name = ?", user).Scan(&stored)
		if err == nil && verify(pass, stored) {
			fmt.Println("Login successful!")
		} else {
			fmt.Println("Invalid credentials.")
		}
	}
}
