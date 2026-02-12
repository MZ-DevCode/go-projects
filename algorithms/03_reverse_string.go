package main

import "fmt"

func reverse(s string) string {
    // 1. Превращаем строку в слайс "рун" (символов Юникода). 
    // Это нужно, чтобы буквы типа 'Я' или 'G' занимали 1 место, а не 2 байта.
    runes := []rune(s)

    // 2. Двойной указатель: i идет с начала (0), j — с конца (длина-1).
    // Цикл работает, пока они не встретятся (i < j).
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        // 3. Магия Go: меняем значения двух переменных местами без третьей переменной!
        runes[i], runes[j] = runes[j], runes[i]
    }
    
    // 4. Превращаем слайс рун обратно в обычную строку
    return string(runes)
}

func main() {
    var text string
    fmt.Print("Text to reverse: ")
    
    // 5. Scan берет текст до первого пробела. Если введешь "Hello World", перевернет только "Hello".
    fmt.Scan(&text)
    
    result := reverse(text)
    fmt.Println("Result: ", result)
}
