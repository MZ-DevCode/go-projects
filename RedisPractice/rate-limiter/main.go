package main

import (
	"context"
	"fmt"
	"time"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	var pass string
	password := "Hello!"

	for {
		fmt.Print("Enter password: ")
		fmt.Scan(&pass)

		val, err := rdb.Get(ctx, "bad_attempts:user1").Int()
		if err != nil && err != redis.Nil {
			fmt.Println("Error:", err)
			return
		}

		if val >= 3 {
			fmt.Println("Access denied. Wait 5 second")
			rdb.Expire(ctx, "bad_attempts:user1", 5*time.Second)
			time.Sleep(5 * time.Second)
			continue
		}

		if pass != password {
			fmt.Println("Wrong password!")
			rdb.Incr(ctx, "bad_attempts:user1")
			rdb.Expire(ctx, "bad_attempts:user1", 5*time.Second)
		} else {
			fmt.Println("Welcome!")
			rdb.Del(ctx, "bad_attempts:user1")
			return
		}
	}
}
