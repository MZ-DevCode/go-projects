package main

import "fmt"

func FindMax(numbers []int) int { 
    // Если слайс пустой — возвращаем 0 (защита от паники)
    if len(numbers) == 0 { 
        return 0
    }
    
    // Берем первый элемент как эталон для сравнения
    max := numbers[0] 
    
    // Цикл по всем числам: "_" (индекс нам не нужен), "val" (текущее число)
    for _, val := range numbers { 
        // Если нашли число покрупнее — обновляем рекорд
        if val > max {
            max = val
        }
    }
    // Когда всех проверили — отдаем чемпиона
    return max
}

func main() {
    list := []int{3, 10, 1, 45, 12, 8}
    fmt.Println("List:", list)
    fmt.Println("Max value:", FindMax(list))
}
