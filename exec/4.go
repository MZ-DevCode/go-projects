package main

import(
	"fmt"
	"os/exec"
	"time"
)

func main(){
	currentTime := time.Now().Format("2006-01-02-_15-04-05")

	for i := 0; i < 2; i++{

	folderName := fmt.Sprintf("backup-%s-%d", currentTime, i)
		cmd := exec.Command("mkdir", folderName)
	
	_ = cmd.Run()	
}}
