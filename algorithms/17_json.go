package main

import(
	"encoding/json"
	"fmt"
)

type Message struct {
	User	string `json:"user"`
	Text	string `json:"message"`
	Status	string `json:"code"`
}
func main() {
	myPost := Message{
		User: "John",
		Text: "Hello, World!",
		Status: "200",
	}
	jsonData, _ := json.Marshal(myPost)
	fmt.Println(string(jsonData))
}
