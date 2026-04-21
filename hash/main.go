package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		http.Error(w, "Error calculating hash", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "SHA-256: %x\n", hash.Sum(nil))
}

func main() {
	http.HandleFunc("/upload", uploadHandler)

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Fatal: %v\n", err)
	}
}
