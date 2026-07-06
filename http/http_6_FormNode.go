package main

import (
	"fmt"
	"net/http"
)

var colorHistory []string

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", mainPage)
	mux.HandleFunc("POST /paint", paintPage)
	http.ListenAndServe(":8080", mux)
}

func mainPage(w http.ResponseWriter, r *http.Request) {

	bg := "white"
	if len(colorHistory) > 0 {
		bg = colorHistory[len(colorHistory)-1]
	}
	fmt.Fprintf(w, `<body style="background-color: %s; min-height: 100vh; margin: 0;">`, bg)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
			<form action="/paint" method="POST">
				<select name="color">
					<option value="red">RED</option>
					<option value="green">GREEN</option>
					<option value="blue">BLUE</option>
				</select>
				<button type="submit">Change color</button>
			</form>
		`)

	for _, value := range colorHistory {
		fmt.Fprintf(w, "<li>%s</li>", value)
	}

}

func paintPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	color := r.FormValue("color")
	colorHistory = append(colorHistory, color)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
