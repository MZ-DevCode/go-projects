package main // Defines this file as an executable program

import (
	"fmt"      // Standard package for formatted I/O (printing to console)
	"net/http" // Package for HTTP client and server implementations
	"strings"  // Package for string manipulation
	"time"     // Package for measuring and displaying time
)

// getHeader is a helper function to extract a specific header from the response.
// It returns "Not specified" if the header is missing to keep the output clean.
func getHeader(r *http.Response, name string) string {
	val := r.Header.Get(name)
	if val == "" {
		return "Not specified"
	}
	return val
}

func main() {
	var url string
	fmt.Print("Enter target URL: ") // Prompt user for input
	fmt.Scan(&url)                  // Read the user's input into the url variable

	// Auto-fix: if the user didn't provide a protocol, default to https
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	// Create a custom HTTP client with a 5-second timeout to prevent hanging
	client := http.Client{Timeout: 5 * time.Second}
	
	// Perform the GET request
	resp, err := client.Get(url)
	if err != nil {
		// If there is a DNS issue or the site is down, print the error and exit
		fmt.Printf("Connection error: %v\n", err)
		return
	}
	
	// defer ensures the response body is closed when main() finishes, 
	// which prevents memory and network connection leaks.
	defer resp.Body.Close()

	// Print the basic connection results
	fmt.Printf("\n[+] Target: %s\n", url)
	fmt.Printf("Status: %d\n", resp.StatusCode) // HTTP status code (e.g., 200, 404)

	// Fetch and display specific headers related to technology and security
	fmt.Printf("Server: %s\n", getHeader(resp, "Server"))
	fmt.Printf("Powered By: %s\n", getHeader(resp, "X-Powered-By"))
	fmt.Printf("CSP Security: %s\n", getHeader(resp, "Content-Security-Policy"))
	fmt.Printf("HSTS: %s\n", getHeader(resp, "Strict-Transport-Security"))
	fmt.Printf("Cookies: %s\n", getHeader(resp, "Set-Cookie"))
}
