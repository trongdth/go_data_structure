package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var works int
	works = 100

	c := random(works)

	// FAN OUT
	w1 := calculate(c)
	w2 := calculate(c)

	// FAN IN
	for out := range merge(w1, w2) {
		fmt.Println(out)
	}
}

func random(w int) <-chan int {
	c := make(chan int)
	go func() {
		for index := 0; index < w; index++ {
			c <- rand.Intn(100)
		}
		close(c)
	}()
	return c
}

func calculate(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v + rand.Intn(2)
		}
		close(out)
	}()
	return out
}

func merge(ins ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(ins))

	for _, v := range ins {
		fmt.Printf("TYPE %T \n", v)
		go func(ch <-chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
