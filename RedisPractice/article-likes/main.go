package main

import(
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
)

var(
	rdb *redis.Client
	ctx = context.Background()
)

func main(){
	mux := http.NewServeMux()

	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost: 6379",
	})

	mux.HandleFunc("/like", func(w http.ResponseWriter, r *http.Request){
		id := r.URL.Query().Get("id")
		if id == ""{
			http.Error(w, "input id article", 400)
			return
		}

		val, err := rdb.Incr(ctx, "likes:" + id).Result()
		if err != nil{
			http.Error(w, "Redis Error: " + err.Error(), 500)
			return
		}

		fmt.Fprintf(w, "Article id %s get view. All views: %d", id, val)
	})

	fmt.Println("Server running")
	http.ListenAndServe(":8080", mux)
}
