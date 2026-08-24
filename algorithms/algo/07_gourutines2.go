package main

import(
	"fmt"
	"time"
)

func sayHello(){
	fmt.Println("Gourutine")
}

func main(){
	go sayHello()
	fmt.Println("Main")
	time.Sleep(time.Millisecond * 10)
}
