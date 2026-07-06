package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"Message"`
	Status  int    `json:"status"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{name}", userHandler)

	http.ListenAndServe(":8080", mux)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	var resp Response
	status := http.StatusOK
	name := r.PathValue("name")
	switch name {
	case "admin":
		resp = Response{Message: "Добро пожаловать, босс!", Status: 200}
	case "hacker":
		status = http.StatusForbidden
		resp = Response{Message: "Сюда нельзя", Status: 403}
	default:
		status = http.StatusNotFound
		resp = Response{Message: "Ошибка: Пользователь " + name + " не найден.", Status: 404}
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.WriteHeader(status)

	json.NewEncoder(w).Encode(resp)
}
