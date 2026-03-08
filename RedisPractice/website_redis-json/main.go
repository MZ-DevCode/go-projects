package main

import(
	"fmt"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"context"
	"strconv"
	"net/http"
)

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func main(){
	mux := http.NewServeMux()
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintln(w, `
				<form action="/save" method="POST">
					<input type="text" name="user_name" placeholder="Name" required>
					<input type="number" name="user_age" placeholder="Age" required>
					<button type="submit">Save User</button>
				</form>
			`)
		})

	mux.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost{
			name := r.FormValue("user_name")
			ageInt := r.FormValue("user_age")
			age, err := strconv.Atoi(ageInt)
			if err != nil {
				http.Error(w, "Invalid age", http.StatusBadRequest)
				return
			}

			user := User{Name: name, Age: age}
			jsonData, err := json.Marshal(user)
			err = rdb.Set(ctx, "user:1", jsonData, 0).Err()
			if err != nil {
				http.Error(w, "Error", http.StatusInternalServerError)
				return
		}
			get, err := rdb.Get(ctx, "user:1").Result()
			if err != nil {
				http.Error(w, "Error", http.StatusInternalServerError)
				return
			}
			fmt.Fprintln(w, "User saved successfully!", get)
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
