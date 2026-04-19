package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct{
	ID	int `json:"id"`
	Content string `json:"content"`
	Done	bool `json:"done"`
}

var tasks = []Task{
	{ID: 1, Content: "Buy groceries", Done: false},
}
func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Server", "Go server")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	})

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request){
              	w.Header().Set("Server", "Go server")
		val := r.URL.Query().Get("text")
		newTask := Task{
			ID: len(tasks) + 1,
			Content: val,
			Done: false,
		}
		tasks = append(tasks, newTask)
		fmt.Fprint(w, "Задание добавлено")
	})

	mux.HandleFunc("/clear", func(w http.ResponseWriter, r *http.Request){
		tasks = nil
		fmt.Fprint(w, "Задания очищены")
	})
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", mux)

}
