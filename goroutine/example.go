package main

import (
	"fmt"
)

func doSomething(x int) int {
	return x * 5
}

func main() {
	//channel
	ch := make(chan int)
	//goroutine that return value to main goroutine through a channel
	go func() {
		ch <- doSomething(6)
	}()
	//get value from channel
	fmt.Println(<-ch)
}
