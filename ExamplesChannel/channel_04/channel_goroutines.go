package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	go sendgoroutines(10, c)
	printChannel(c)
}

func sendgoroutines(goroutines int, c chan<- int) {
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go fillchannel(3, c)
	}
	wg.Wait()
	close(c)
}

func fillchannel(nums int, c chan<- int) {
	for i := 0; i < nums; i++ {
		c <- i
	}
	wg.Done()
}

func printChannel(c <-chan int) {
	for v := range c {
		fmt.Println("val ", v)
	}
}
