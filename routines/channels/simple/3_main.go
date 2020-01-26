package main

import "fmt"

func main() {
	c := incrementor()
	csum := sum(c)
	for sum := range csum {
		fmt.Println(sum)
	}
}

func incrementor() chan int {
	c := make(chan int)
	go func() {
		for index := 0; index < 10; index++ {
			c <- index
		}
		close(c)
	}()
	return c
}

func sum(in chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
