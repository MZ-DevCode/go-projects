package main // Defines this file as an executable program

import (
	"fmt"      // Package for formatted I/O (printing and string formatting)
	"net/http" // Package for making HTTP requests
	"strings"  // Package for string manipulation (prefix checks)
	"os"       // Package for operating system functions (file handling)
	"time"     // Package for handling time and timeouts
)

// getHeader is a helper function to safely extract headers.
// It returns "Not specified" if the header is missing to keep the report clean.
func getHeader(r *http.Response, name string) string {
	val := r.Header.Get(name)
	if val == "" {
		return "Not specified"
	}
	return val
}

func main() {
	var url string
	fmt.Print("Enter target URL: ") // Prompt user for the website address
	fmt.Scan(&url)                  // Read user input into the 'url' variable

	// Logic: if the user didn't type http/https, we add it automatically
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	// Create an HTTP client with a 5-second limit to prevent the app from hanging
	client := http.Client{Timeout: 5 * time.Second}
	
	// Send the actual GET request to the server
	resp, err := client.Get(url)
	if err != nil {
		// If there's a DNS error, timeout, or no internet, print it and stop
		fmt.Printf("Connection error: %v\n", err)
		return
	}
	
	// 'defer' ensures the connection is closed when the function finishes.
	// This is critical to prevent memory and network leaks.
	defer resp.Body.Close()

	// Use Sprintf to format data into strings instead of printing them immediately.
	// This allows us to store the results in variables.
	target := fmt.Sprintf("\n[+] Target: %s\n", url)
	status := fmt.Sprintf("Status: %d\n", resp.StatusCode)

	// Extracting specific security and technology headers
	server := fmt.Sprintf("Server: %s\n", getHeader(resp, "Server"))
	powered := fmt.Sprintf("Powered By: %s\n", getHeader(resp, "X-Powered-By"))
	csp := fmt.Sprintf("CSP Security: %s\n", getHeader(resp, "Content-Security-Policy"))
	hsts := fmt.Sprintf("HSTS: %s\n", getHeader(resp, "Strict-Transport-Security"))
	cookies := fmt.Sprintf("Cookies: %s\n", getHeader(resp, "Set-Cookie"))

	// Print everything to the terminal for the user to see
	fmt.Print(target, status, server, powered, csp, hsts, cookies)

	// Combine all separate strings into one large 'text' block using concatenation (+)
	text := target + status + server + powered + csp + hsts + cookies

	// Open or create "scan.txt"
	// O_APPEND: add to the end of the file; O_CREATE: make file if it doesn't exist; O_WRONLY: write mode
	file, err := os.OpenFile("scan.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is saved and closed properly

	// Write the entire text block into the file in one operation
	file.WriteString(text)
	fmt.Println("Saved to scan.txt")
}
