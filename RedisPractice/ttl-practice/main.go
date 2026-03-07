package main

import(
	"context"
	"fmt"
	"time"
	"github.com/redis/go-redis/v9"
)

var(
	rdb *redis.Client
	ctx = context.Background()
)

func main(){
	rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DialTimeout:  2 * time.Second,
	})

	err := rdb.Ping(ctx).Err()
	if err != nil{
		fmt.Println("Error: ", err)
		return
	}

	err = rdb.Set(ctx, "user:1", "Alex", 5 * time.Second).Err()
	if err != nil{
		fmt.Println("Error: ", err)
	}
	fmt.Println("recorded to redis")
	val, err := rdb.Get(ctx, "user:1").Result()
	if err != nil{
		fmt.Println("Error: ", err)
	}
	fmt.Printf("Read: %s", val)
	fmt.Println("Wait 6 second")
	for i := 0; i < 6; i++{
		fmt.Println(i + 1)
		time.Sleep(time.Second)
	}

	val2, err := rdb.Get(ctx, "user:1").Result()
	if err == redis.Nil{
		fmt.Println("Redis deleted key")
	} else if err != nil{
		fmt.Printf("Error: %v", err)
	} else{
		fmt.Printf("Key still here: %s", val2)
	}
}
