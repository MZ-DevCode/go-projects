package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func getHeader(r *http.Response, name string) string {
	val := r.Header.Get(name)
	if val == "" {
		return "Not specified"
	}
	return val
}

func main() {
	var url string
	fmt.Print("Enter target URL: ")
	fmt.Scan(&url)

	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Connection error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("\n[+] Target: %s\n", url)
	fmt.Printf("Status: %d\n", resp.StatusCode)

	fmt.Printf("Server: %s\n", getHeader(resp, "Server"))
	fmt.Printf("Powered By: %s\n", getHeader(resp, "X-Powered-By"))
	fmt.Printf("CSP Security: %s\n", getHeader(resp, "Content-Security-Policy"))
	fmt.Printf("HSTS: %s\n", getHeader(resp, "Strict-Transport-Security"))
	fmt.Printf("Cookies: %s\n", getHeader(resp, "Set-Cookie"))
}
