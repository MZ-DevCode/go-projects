package main

import (
	"fmt"
	"net/http"
	"net"
	"strings"
	"os"
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

	clean := strings.Replace(url, "https://", "", 1)
	clean = strings.Replace(clean, "http://", "", 1)
	clean = strings.TrimSuffix(clean, "/")

	ips, err := net.LookupHost(clean)

	target := fmt.Sprintf("\n[+] Target: %s\n", url)
	ip := fmt.Sprintf("[*] IP: %v\n", ips)
	status := fmt.Sprintf("Status: %d\n", resp.StatusCode)
	server := fmt.Sprintf("Server: %s\n", getHeader(resp, "Server"))
	powered := fmt.Sprintf("Powered By: %s\n", getHeader(resp, "X-Powered-By"))
	csp := fmt.Sprintf("CSP Security: %s\n", getHeader(resp, "Content-Security-Policy"))
	hsts := fmt.Sprintf("HSTS: %s\n", getHeader(resp, "Strict-Transport-Security"))
	cookies := fmt.Sprintf("Cookies: %s\n", getHeader(resp, "Set-Cookie"))

	fmt.Print(target, ip, status, server, powered, csp, hsts, cookies)

	text := target + ip + status + server + powered + csp + hsts + cookies


	file, err := os.OpenFile("scan.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	file.WriteString(text)
	fmt.Println("Saved to scan.txt")
}
