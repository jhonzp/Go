package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	fanin := make(chan int)

	go enviar(par, impar)

	go recibir(par, impar, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Finalizando.")
}

// send channel
func enviar(par, impar chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
}

// receive channel
func recibir(par, impar <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range par {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range impar {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
