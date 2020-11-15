package main

import (
	"fmt"

	"github.com/trongdth/go_practise/array"
)

func main() {
	output := array.LongestPeak([]int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3})
	fmt.Println(output)
}
