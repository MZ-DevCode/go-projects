package main

import(
	"fmt"
	"os/exec"
)

func main(){
	cmd := exec.Command("echo", "hello world")
	result, _ := cmd.Output()
	fmt.Print(string(result))
}


