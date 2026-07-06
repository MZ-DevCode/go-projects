package main

import (
	"fmt"
	"sync"
	"time"
)

// Добавили возвращаемый тип string в конце сигнатуры
func downloadFile(filename string, size int, wg *sync.WaitGroup) string {
	if wg != nil {
		defer wg.Done()
	}

	fmt.Printf("Началось скачивание файла %s...\n", filename)
	time.Sleep(time.Duration(size) * time.Millisecond)
	
	// Теперь функция возвращает строку с результатом работы
	return fmt.Sprintf("Файл %s успешно скачан!", filename)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	// В main мы по-прежнему запускаем их как горутины
	go downloadFile("document.pdf", 100, &wg)
	go downloadFile("music.mp3", 300, &wg)
	go downloadFile("video.mp4", 800, &wg)

	wg.Wait()
	fmt.Println("Все файлы загружены. Программа завершена!")
}

