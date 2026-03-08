package main

import(
	"fmt"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"context"
)

var ctx = context.Background()

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func main(){
	var name string
	var age int
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	user := User{Name: name, Age: age}

	jsonData, err := json.Marshal(user)
	if err != nil{
		fmt.Println("Error: ", err)
	}

	err = rdb.Set(ctx, "user:1", jsonData, 0).Err()
	if err != nil{
		fmt.Println("Error: ", err)
	}

	val, err := rdb.Get(ctx, "user:1").Result()
	if err != nil{
		fmt.Println("Error: ", err)
	}

	fmt.Println("JSON from Redis:", val)
}
