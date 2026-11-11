package main

import "net/http"

func renderHTML() {
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
			<input type="text"></input>
			<input name="add" type="submit">Add task</input>
		</body>
		</html>
	`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderHTML)
	http.ListenAndServe(":8080",mux
}
