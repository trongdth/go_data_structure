package main

import (
	"log"

	"github.com/trongdth/go_practise/array"
)

func main() {
	output := array.IsMonotonic([]int{-1, 5, 10, 20, 28, 3})
	log.Println(output)
}
