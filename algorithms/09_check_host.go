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

	client := http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Connection error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("\n[+] Target: %s\n", url)
	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Server: %s\n", resp.Header.Get("Server"))
	fmt.Printf("Powered By: %s\n", resp.Header.Get("X-Powered-By"))
	fmt.Printf("CSP Security: %s\n", resp.Header.Get("Content-Security-Policy"))
	fmt.Printf("HSTS (HTTPS Only): %s\n", resp.Header.Get("Strict-Transport-Security"))
	fmt.Printf("Cookies: %s\n", resp.Header.Get("Set-Cookie"))
}
