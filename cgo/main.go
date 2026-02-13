package main

/*
#include "fast_logic.h"
*/
import "C"
import "fmt"

func main() {
    nums := []int32{1, 2, 3, 4, 5}

    C.process_data((*C.int)(&nums[0]), C.int(len(nums)))

    fmt.Println(nums)
}
