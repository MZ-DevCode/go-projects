package main

import(
	"fmt"
	"github.com/redis/go-redis/v9"
	"encoding/json"
	"context"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func main(){
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	dsn := "root:123@tcp(127.0.0.1:3306)/test_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil{
		fmt.Println("Error conect to MySQL", err)
	}

	defer db.Close()

	userID := 1
	key := fmt.Sprintf("user:%d", userID)

	var user User
	val, err := rdb.Get(ctx, key).Result()
	if err == nil{
		fmt.Println("User found in Redis:", val)
		json.Unmarshal([]byte(val), &user)
	} else {
		fmt.Println("User not found in Redis, querying MySQL...")
		err = db.QueryRow("Select id, name, email FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Name, &user.Email)
		if err != nil{
			fmt.Println("Error querying MySQL:", err)
		}
		data, _ := json.Marshal(user)
		rdb.Set(ctx, key, data, 10 * time.Minute)
		fmt.Println("User found in MySQL and cached in Redis:", user)
	}

	finalJson, _ := json.Marshal(user)
	fmt.Println("Final User Data:", string(finalJson))
}
