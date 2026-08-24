package main

import "fmt"

func main(){
var number int = 345

ones := number % 10

withoutLast := number / 10

tens := withoutLast % 10

hundreds := withoutLast / 10

    fmt.Printf("Исходное число: %d\n", number)
    fmt.Printf("Сотни: %d\n", hundreds)
    fmt.Printf("Десятки: %d\n", tens)
    fmt.Printf("Единицы: %d\n", ones)


}
