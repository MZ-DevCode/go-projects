package main

import "fmt"

func isPalindrome(s string) bool {
    // Превращаем в руны, чтобы корректно работало с кириллицей (Юникод)
    runes := []rune(s)
    
    // Два указателя: i с начала, j с конца. Сближаем их к середине.
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        // Если буквы на краях РАЗНЫЕ — это точно не палиндром
        if runes[i] != runes[j] {
            return false // Выходим из функции немедленно
        }
    }
    // Если цикл дошел до конца и расхождений не нашел — всё честно
    return true
}

func main() {
    var word string
    fmt.Print("Print word: ")
    // Читаем одно слово из терминала
    fmt.Scan(&word)

    // Проверяем результат функции
    if isPalindrome(word) {
        fmt.Println("It's palindrome")
    } else {
        fmt.Println("It's not palindrome")
    }
}
