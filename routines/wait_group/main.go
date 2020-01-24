package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg      sync.WaitGroup
	mutex   sync.Mutex
	counter int
	grs     int
)

// this is used for setting the parallelism (<= 1.5)
func init() {
	grs = 2
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(grs)
	for index := 0; index < grs; index++ {
		go increament(index)
	}
	wg.Wait()

	fmt.Println("Final counter : ", counter)
}

func increament(i int) {
	for index := 0; index < 10; index++ {
		mutex.Lock()
		fmt.Println("go routine : ", i, " - counter = ", counter)
		counter++
		mutex.Unlock()
	}
	wg.Done()
}
