package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)

	mux.HandleFunc("GET /profile/{name}", profileHandler)

	mux.HandleFunc("GET /status", statusHandler)

	fmt.Println("Сервер запущен на :8080...")
	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Добро пожаловать на главную страницу!")
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	fmt.Fprintf(w, "Это профиль пользователя: %s", name)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Server-Status", "Online")
	fmt.Fprint(w, "Сервер работает нормально")
}
