package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for index := 0; index < 10; index++ {
			c <- index
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
