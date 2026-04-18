package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
)

var rdb *redis.Client
var ctx = context.Background()

func main(){
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	err := rdb.Set(ctx, "test", "hello redis", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		val, err := rdb.Get(ctx, "test").Result()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, "Hello, %s", val)
	})

	fmt.Print("Server running")
	http.ListenAndServe(":8080", nil)

}
