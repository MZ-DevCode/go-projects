package main

import (
	"context"
	"fmt"
	"net/http"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr: "db:6379",
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		hits, err := rdb.Incr(ctx, "hits").Result()
		if err != nil {
			http.Error(w, "Redis error", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "You have visited this page %d times.", hits)
	})

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request){
		_, err := rdb.Del(ctx, "hits").Result()
		if err != nil{
			http.Error(w, "Redis error", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Counter reset.")
	})

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
