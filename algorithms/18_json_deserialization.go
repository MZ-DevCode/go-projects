package main

import(
	"encoding/json"
	"fmt"
)

type UserReg struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password int `json:"password"`
}

func main() {
	jsonText := `{"username":"john", "email":"", "password":123}`
	var user UserReg
	json.Unmarshal([]byte(jsonText), &user)
	if user.Email == "" {
		fmt.Println("Please provide an email address.")
	} else{
		fmt.Printf("User %s is valid", user.Username)
	}
}
