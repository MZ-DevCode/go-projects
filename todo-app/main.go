package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tasks := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- TODO LIST ---")
		fmt.Println("1. Show tasks")
		fmt.Println("2. Add task")
		fmt.Println("3. Delete task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an action (1-4): ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			if len(tasks) == 0 {
				fmt.Println("The list is currently empty!")
			} else {
				fmt.Println("Your Tasks:")
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}
		case "2":
			fmt.Print("Enter task description: ")
			scanner.Scan()
			text := scanner.Text()
			if text != "" {
				tasks = append(tasks, text)
				fmt.Println("Task added!")
			} else {
				fmt.Println("Error: Task description cannot be empty!")
			}
		case "3":
			if len(tasks) == 0 {
				fmt.Println("The list is empty!")
			} else {
				fmt.Print("Enter task number to delete: ")
				scanner.Scan()
				fmt.Println("Feature coming soon...")
			}
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
