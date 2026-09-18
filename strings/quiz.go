package main

import(
	"fmt"
	"strings"
	"bufio"
	"strconv"
	"os"
)

func main(){
	fmt.Println("Квиз\n")
	score := 0
	total := 3

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Столица Польши: ")
	input1, _ := reader.ReadString('\n')
	answer1 := strings.TrimSpace(strings.ToLower(input1))

	switch answer1{
	  case "варшава":
		fmt.Println("Правильно")
		score++
	  default:
		fmt.Println("Неправильно")

	}

	fmt.Print("2+2: ")
        input2, _ := reader.ReadString('\n')
        xz := strings.TrimSpace(input2)
	answer2, _ := strconv.Atoi(xz)
	if answer2 != 4{
		fmt.Println("Не правильно")
	} else{
		fmt.Println("Правильно")
                score++
	}		

	 fmt.Print("Море над Польшей: ")
        input3, _ := reader.ReadString('\n')
        answer3 := strings.TrimSpace(strings.ToLower(input3))

        switch answer3{
          case "балтийское":
                fmt.Println("Правильно")
                score++
          default:
                fmt.Println("Неправильно")
        }

	percent := ((float64(score) / float64(total)) * 100)

	fmt.Printf("Набрано очков: %d из %d\n", score, total)
	fmt.Printf("Результат: %.1f%%\n", percent)
	
}
