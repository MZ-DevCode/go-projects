package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("mkdir", "my_secret_folder")

	err := cmd.Run()

	if err != nil {
		fmt.Println("Не удалось создать папку:", err)
		return
	}

	fmt.Println("Папка успешно создана!")
}
