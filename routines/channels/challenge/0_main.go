package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func main() {
	o := sum(random(), random())

	var sum int
	for n := range o {
		sum += n
	}

	fmt.Println(sum)

}

func random() <-chan int {
	c := make(chan int)
	go func() {
		c <- rand.Intn(30)
	}()
	return c
}

// FAN IN
func sum(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	wg.Add(2)

	go func() {
		out <- <-in1
		wg.Done()
	}()

	go func() {
		out <- <-in2
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
