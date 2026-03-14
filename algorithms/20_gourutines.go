package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup){
	defer wg.Done()
	for i := 0; i < 5; i++{
		fmt.Println(i+1, " - ", text)
	}
}

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go say("I work in a goroutine", &wg)
	wg.Wait()
	fmt.Println("All tasks are done")
}
