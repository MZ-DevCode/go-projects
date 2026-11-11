package main

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
			</ul>
			<input type="text"></input>
			<input name="add" type="submit">Add task</input>
		</body>
		</html>
	`
}
