package main

import "fmt"

func isPalindrome(s string) bool {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        if runes[i] != runes[j] {
            return false
        }
    }
    return true
}

func main() {
    var word string
    fmt.Print("Print word: ")
    fmt.Scan(&word)

    if isPalindrome(word) {
        fmt.Println("It's pallindrome")
    } else {
        fmt.Println("It's no pallindrome")
    }
}
