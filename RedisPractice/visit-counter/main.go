package main

import(
	"context"
	"database/sql"
	"fmt"
	"net/http"
	_"github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var(
	db *sql.DB
	rdb *redis.Client
	ctx = context.Background()
)

func main(){
	mux := http.NewServeMux()
	var err error
	db, err = sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/todo_db")
	if err != nil{
		fmt.Println("Error: ", err)
		return
	}
	rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	views, err := rdb.Incr(ctx, "main_page_views").Result()
	if err != nil{
		http.Error(w, "Redis Error: ", err)
	}
	fmt.Fprintf(w, "<p>This page has been viewed %d times (data from Redis)</p>", views)
	fmt.Fprint(w, "<a href='/sync'><button>Save to database</button>")
	})

	mux.HandleFunc("/sync", func(w http.ResponseWriter, r *http.Request){
	views, err := rdb.Get(ctx, "main_page_views").Int()
	if err != nil{
	http.Error(w, "Error: ", 500)
	return
	}

	_, err = db.Exec("INSERT INTO page_stats (page_name, views) VALUES ('main', ?) ON DUPLICATE KEY UPDATE views = ?", views, views)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprint(w, "Data saved <a href='/'>Back</a>")
	})
	fmt.Println("Server running")
	http.ListenAndServe(":8080", mux)
}
