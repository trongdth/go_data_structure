package main

import (
	"fmt"

	"github.com/trongdth/go_practise/array"
)

func main() {
	output := array.SpiralTraverse([][]int{{1, 2, 3, 4}, {10, 11, 12, 5}, {9, 8, 7, 6}})
	fmt.Println(output)
}