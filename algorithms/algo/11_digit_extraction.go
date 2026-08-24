package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Print("Print your number: ")
	fmt.Scan(&n)
	summa := 0

	for n > 0{
		digit := n % 10
		summa += digit
		n = n / 10
	}

	fmt.Println(summa)
}
