package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func gatewayHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8081/users")
	if err != nil {
		http.Error(w, "Сервис недоступен", http.StatusServiceUnavailable)
	}
	defer resp.Body.Close()

	var users interface{}
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		http.Error(w, "Ошибка данных", http.StatusInternalServerError)
		return
	}

	// Сразу кодируем их в ответ нашего Gateway
	w.Header().Set("Content-Type", "text/html")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/get-users", gatewayHandler)
	fmt.Println("Gateway на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
