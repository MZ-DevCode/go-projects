package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintln(w, `
			<h1>Redis + JSON Management System</h1>

			<div style="border: 1px solid #ccc; padding: 10px; margin-bottom: 10px;">
				<h3>1. Create (Save to Redis)</h3>
				<form action="/save" method="POST">
					<input type="text" name="user_name" placeholder="Name" required>
					<input type="number" name="user_age" placeholder="Age" required>
					<button type="submit">Save</button>
				</form>
			</div>

			<div style="border: 1px solid #ccc; padding: 10px; margin-bottom: 10px;">
				<h3>2. Find (Read + Unmarshal)</h3>
				<form action="/get" method="GET">
					<input type="text" name="name" placeholder="Enter name to search">
					<button type="submit">Search</button>
				</form>
			</div>

			<div style="border: 1px solid #ccc; padding: 10px; margin-bottom: 10px;">
				<h3>3. UPDATE (Edit existing user)</h3>
				<form action="/update" method="GET">
					<input type="text" name="name" placeholder="Target Name" required>
					<input type="number" name="age" placeholder="New Age" required>
					<button type="submit">Update</button>
				</form>
			</div>

			<div style="border: 1px solid #ccc; padding: 10px;">
				<h3>4. Delete (DEL Command)</h3>
				<form action="/delete" method="POST">
					<input type="text" name="name" placeholder="Name to delete">
					<button type="submit">Delete</button>
				</form>
			</div>
		`)
	})

	mux.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		name := r.FormValue("user_name")
		ageStr := r.FormValue("user_age")
		age, _ := strconv.Atoi(ageStr)

		user := User{Name: name, Age: age}
		jsonData, _ := json.Marshal(user)

		err := rdb.Set(ctx, "user:"+name, jsonData, 0).Err()
		if err != nil {
			http.Error(w, "Redis Error", 500)
			return
		}
		fmt.Fprintf(w, "User %s successfully saved! <br><a href='/'>Back</a>", name)
	})

	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		key := "user:" + name
		val, err := rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			fmt.Fprintf(w, "User %s not found! <br><a href='/'>Back</a>", name)
			return
		} else if err != nil {
			http.Error(w, "Redis Error", 500)
			return
		}
		var u User
		json.Unmarshal([]byte(val), &u)
		fmt.Fprintf(w, "<h1>Found in Redis:</h1> <p>Name: %s</p> <p>Age: %d</p> <a href='/'>Back</a>", u.Name, u.Age)
	})

	mux.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		ageStr := r.URL.Query().Get("age")
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			http.Error(w, "Invalid age value", http.StatusBadRequest)
			return
		}

		key := "user:" + name
		exists, _ := rdb.Exists(ctx, key).Result()
		if exists == 0 {
			fmt.Fprintf(w, "User %s does not exist. Create it first! <br><a href='/'>Back</a>", name)
			return
		}

		user := User{Name: name, Age: age}
		jsonData, _ := json.Marshal(user)
		rdb.Set(ctx, key, jsonData, 0)

		fmt.Fprintf(w, "User %s updated to age %d! <br><a href='/'>Back</a>", name, age)
	})

	mux.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		key := "user:" + name
		deleted, _ := rdb.Del(ctx, key).Result()
		if deleted == 0 {
			fmt.Fprintf(w, "Key %s not found. <br><a href='/'>Back</a>", key)
			return
		}
		fmt.Fprintf(w, "User %s deleted! <br><a href='/'>Back</a>", name)
	})

	fmt.Println("Server is ready at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
