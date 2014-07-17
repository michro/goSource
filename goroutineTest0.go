package main

import (
	"fmt"
	// "time"
	"runtime"
)

func main() {
	runtime.Gosched()
	go Go()
	// time.Sleep(2 * time.Second)
}

func Go() {
	fmt.Println("go go go!")
}
