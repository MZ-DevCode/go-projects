package main

import("fmt")

func MinNumber(minNumber []int) int{
	min := minNumber[0]
	for _, val := range minNumber{
		if val < min{
			min = val
		}
}
return min
}

func main(){
		list := []int{15, 2, 45, -5, 10}
		fmt.Println("Min number: ", MinNumber(list))
	}
