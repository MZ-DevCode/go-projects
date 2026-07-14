package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Alex", Email: "alex@example.com"},
	{ID: 2, Name: "Maria", Email: "maria@example.com"},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", usersHandler)
	fmt.Println("User Service успешно запущен на http://localhost:8081")
	http.ListenAndServe(":8081", mux)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	json.NewEncoder(w).Encode(users)
}
