package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(i+1, "-", text)
	}
}

func test(a, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	res := a + b
	fmt.Println("Результат теста:", res)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go say("goroutines", &wg)
	go test(5, 10, &wg)

	wg.Wait()

	fmt.Println("Все задачи выполнены, main завершается")
}
