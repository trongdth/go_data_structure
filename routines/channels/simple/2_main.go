package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for index := 0; index < 10; index++ {
			c <- index
		}
		done <- true
	}()

	go func() {
		for index := 0; index < 10; index++ {
			c <- index
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
