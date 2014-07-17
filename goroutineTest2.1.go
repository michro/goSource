package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	wg.Wait()
}

func Go(wg *sync.WaitGroup, index int) {
	var a int64 = 1
	for i := 0; i < 1000000; i++ {
		a += int64(i)
	}
	fmt.Println(index, a)
	wg.Done()
}
