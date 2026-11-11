package main

import (
	"log"
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
		taskValue := r.FormValue("task")
		var tasks string = append(tasks, Data{
			Task: taskValue
		})

		tmpl, err := template.New("todo").Parse(htmlTemplate)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}

		err = tmpl.Execute(w, tasks)
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderHTML)
	err := http.ListenAndServe(":8080", mux)
	if err != nil{
		log.Fatal(err)
	}
}
