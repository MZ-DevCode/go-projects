package main

import(
	"os/exec"
	"fmt"
)

func main(){
	cmd := exec.Command("git", "--version")
	result, _ := cmd.Output()
	fmt.Print(string(result))
}
