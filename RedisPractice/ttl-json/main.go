package main

import(
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"context"
	"time"
)

type User struct{
	UserID int `json:"user_id"`
	Token string `json:"token"`
}

func main(){
	var userID = 123
	var token = "shjaovuiesafsdf"

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	user := User{UserID: userID, Token: token}
	jsonData, err := json.Marshal(user)
	if err != nil{
		fmt.Println("Error: ", err)
	}

	err = rdb.Set(ctx, "user:123", jsonData, 5 * time.Second).Err()
	if err != nil{
		fmt.Println("Error: ", err)
	}

	val, err := rdb.Get(ctx, "user:123").Result()
	if err != nil{
		fmt.Println("Error: ", err)
	}

	fmt.Println("JSON from Redis:", val)
	fmt.Println("Waiting for 5 seconds...")
	for i := 0; i < 5; i++{
		fmt.Printf("%d seconds...\n", i+1)
		time.Sleep(1 * time.Second)
	}

	val, err = rdb.Get(ctx, "user:123").Result()
	if err == redis.Nil{
		fmt.Println("Key does not exist")
	} else if err != nil{
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("JSON from Redis:", val)
	}
	fmt.Println("Key deleted", val)
}
