package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://www.google.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("inaccessible website", err)
		return
	}

	fmt.Printf("Site: %s\nStatus: %d\n", url, resp.StatusCode)
}
