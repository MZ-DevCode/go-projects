package main

import(
	"encoding/json"
	"fmt"
)

func main() {
	jsonText := `{"name":"John", "age":25}`

	var data map[string]interface{}
	json.Unmarshal([]byte(jsonText), &data)

fmt.Println(data["name"])
}
