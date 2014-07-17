package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go(c, i)
		// <-c
	}
	for j := 0; j < 10; j++ {
		<-c
	}
}

func Go(c chan bool, index int) {
	var a int64 = 1
	for i := 0; i < 1000000; i++ {
		a += int64(i)
	}
	fmt.Println(index, a)

	// if index == 9 {
	c <- true
	// }
}
