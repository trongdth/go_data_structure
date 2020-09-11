package main

import (
	"log"

	"github.com/trongdth/go_practise/array"
)

func main() {
	output := array.SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17})
	log.Println(output)
}
