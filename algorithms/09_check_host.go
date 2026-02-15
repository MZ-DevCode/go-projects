package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var url string

	fmt.Print("Enter target URL: ")
	fmt.Scan(&url)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Connection error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("\n[+] Target: %s\n", url)
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Server Info: %s\n", resp.Header.Get("Server"))
}
