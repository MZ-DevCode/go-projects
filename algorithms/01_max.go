package main

import "fmt"

func FindMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, val := range numbers {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	list := []int{3, 10, 1, 45, 12, 8}
	fmt.Println("List:", list)
	fmt.Println("Max value:", FindMax(list))
}
