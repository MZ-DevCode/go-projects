package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Width(X): ")
	fmt.Scan(&x)
	fmt.Print("Height(Y): ")
	fmt.Scan(&y)

	fmt.Println("\nResult:")

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			fmt.Print("0 ")
		}
		fmt.Println()
	}
}
