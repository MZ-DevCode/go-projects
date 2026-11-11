package main

import (
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Task string
}

var tasks []Data

func renderHTML(w http.ResponseWriter, r *http.Request) {
	var html string = `
		<!DOCTYPE html>
		<html>
		<head>

		</head>
		<body>
			<ul>
				{{ range . }}
				<li>{{ .Task }}</li>
				{{end}}
			</ul>
			<form method="POST" action="/">
				<input type="text" name="task">
				<input name="add" type="submit" value="add">
			</form>
			</body>
		</html>
	`

	if r.Method == http.MethodPost {
		taskValue := r.FormValue("task")
		tasks = append(tasks, Data{
			Task: taskValue,
		})
	}
	tmpl, err := template.New("todo").Parse(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderHTML)
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
