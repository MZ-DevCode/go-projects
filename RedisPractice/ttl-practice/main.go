package main

import (
	"context"
	"fmt"
	"time"
	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		DialTimeout:  2 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	userId := "user:1001"

	err := rdb.HSet(ctx, userId, map[string]interface{}{
		"name": "Alex",
		"age":  25,
	}).Err()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	rdb.Expire(ctx, userId, 5*time.Second)
	fmt.Println("User created (Hash) with 5s TTL")

	newAge, err := rdb.HIncrBy(ctx, userId, "age", 1).Result()
	if err == nil {
		fmt.Printf("In-memory increment: age is now %d\n", newAge)
	}

	user, err := rdb.HGetAll(ctx, userId).Result()
	if err == nil {
		fmt.Printf("Data: Name: %s, Age: %s\n", user["name"], user["age"])
	}

	fmt.Println("Waiting 6 seconds...")
	for i := 1; i <= 6; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Second)
	}
	fmt.Println()

	exists, _ := rdb.Exists(ctx, userId).Result()
	if exists == 0 {
		fmt.Println("Result: Redis auto-deleted data (TTL)")
	} else {
		fmt.Println("Result: Data still exists")
	}
}
