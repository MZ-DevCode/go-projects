package main
import "fmt"

func main() {
  var x int = 10
  var y byte = 100
  var sum1 int = x + int(y)
  var sum2 byte = byte(x) + y
  fmt.Println(sum1, sum2)
}
