package main

import("fmt")

func MinNumber(minNumber []int) (int, int){
	min := minNumber[0]
	minIndex := 0
	for i, val := range minNumber{
		if val < min{
			min = val
			minIndex = i
		}
	        
}
return min, minIndex
}

func main(){
		list := []int{15, 2, 45, -5, 10}
		val, i := MinNumber(list)
		fmt.Printf("Min number: %d, Index: %d", val, i)
	}
