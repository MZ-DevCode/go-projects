package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHello)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	wc, err := fmt.Fprint(w, "Hello World")
	if err != nil {
		slog.Error("Error writing response", "err", err)
		return
	}

	fmt.Printf("%d bytes written\n", wc)
}
