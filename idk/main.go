package main

import (
	"net/http"

	"google.golang.org/api/tasks/v1"
)

type Data struct {
	Task string
}

func renderHTML(w http.ResponseWriter, r *http.Request) {
	var html string = `
		<!DOCTYPE html>
		<html>
		<head>

		</head>
		<body>
			<ul>
				{{ range . }}
				<li>{{ .Tasks }}</li>
				{{end}}
			</ul>
			<input type="text" name="task"></input>
			<input name="add" type="submit">Add task</input>
		</body>
		</html>
	`

	if r.Method == http.MethodPost {
		task := r.FormValue("task")
		tasks = append(tasks, Data{
			Task: task
		})

		err, tmpl := template.New("tmpl").Parse(tasks)
		if err != nil{
			panic(err)
		}

	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderHTML)
	http.ListenAndServe(":8080", mux)
}
