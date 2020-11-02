package main

import (
	"fmt"
)

func main() {
	c := gen()
	recibir(c)

	fmt.Println("A punto de finalizar.")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func recibir(c <-chan int) {
	for v := range c {
		fmt.Println("Valor ", v)
	}
}
