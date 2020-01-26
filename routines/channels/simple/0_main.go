package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for index := 0; index < 10; index++ {
			c <- index
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}

	}()
	time.Sleep(time.Second)
}
