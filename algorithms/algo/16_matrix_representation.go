package main

import "fmt"

func main() {
    var x [2][3]int
    x[0][1] = 5
    x[1][2] = 10
    
    for _, row := range x{
        for _, val := range row{
            fmt.Print(val, " ")
        }
        fmt.Println("")
    }
}
