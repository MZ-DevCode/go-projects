package main

import(
	"database/sql"
	"fmt"
	"net/http"
      _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main(){
	mux := http.NewServeMux()
	db, _ = sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test_db")
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if r.Method == "POST"{
			text := r.FormValue("message")
			db.Exec("INSERT INTO notes VALUES (?)", text)
		}

		fmt.Fprint(w, `
		<form method="POST">
			<input name="message">
			<button>Save to database</button>
		</form>
		<hr>
		`)

		rows, _ := db.Query("SELECT content FROM notes")
		for rows.Next(){
			var msg string
			rows.Scan(&msg)
			fmt.Fprintf(w, "<li>%s</li>", msg)
		}
	})

	fmt.Println("Server running")
	http.ListenAndServe(":8080", mux)
}
