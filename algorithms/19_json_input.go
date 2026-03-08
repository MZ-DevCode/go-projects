package main

import(
	"encoding/json"
	"fmt"
)

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func main(){
	var name string
	var age int
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	user := User{Name: name, Age: age}
	jsonData, err := json.Marshal(user)
	if err != nil{
		fmt.Println("Error: ", err)
	}
	fmt.Println("JSON:", string(jsonData))
}
