package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	start := time.Now()
	for i := 0; i < 100; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
	fmt.Println("Time elapsed:", time.Since(start))
}

func main() {
	say("world")
	say("hello")
	say("!!!!")
}
