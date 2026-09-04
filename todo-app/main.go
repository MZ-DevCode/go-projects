package main

import(
	"fmt"
	"os"
	"strconv"
)

type Task struct{
	ID 	int	`json:"id"`
	Text 	string	`json:"text"`
	Status 	bool	`json:"status"`
}

const fileName = "tasks.json"

func loadTasks() []Task{
	file, err := os.ReadFile(fileName)
	if err != nil{
		log.Println("File doesn't exist: ")
		return []Task{}
	}

	var tasks []Task

	json.Unmarshal(file, &tasks)

	return tasks
}

func main(){
		
}
