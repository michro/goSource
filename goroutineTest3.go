package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	// o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					// o <- true
					fmt.Println("Failed to read c1.")
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					// o <- true
					fmt.Println("Failed to read c2.")
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "hi"
	c1 <- 2
	c2 <- "buddy"

	// close(c1)
	// close(c2)

	// <-o
}
