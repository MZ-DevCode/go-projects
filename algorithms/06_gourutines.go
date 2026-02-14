package main

import (
	"fmt"
	"time"
)

func say(s string, n int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Поток %s: итерация %d\n", s, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go say("Yes", 1)
	say("No", 2)
}
