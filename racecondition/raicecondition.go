package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0

//wait group
var wg sync.WaitGroup

func goroutinefunc() {
	v := counter
	//Yield the procesor
	runtime.Gosched()
	v++
	counter = v

	//was done
	wg.Done()
}

func main() {

	fmt.Println("CPUs", runtime.NumCPU())
	//amount of goroutine must be wait
	wg.Add(6)
	go goroutinefunc()
	go goroutinefunc()
	go goroutinefunc()
	go goroutinefunc()
	go goroutinefunc()
	go goroutinefunc()
	//while all goroutime are finished
	wg.Wait()
	fmt.Println(counter)
}

//run this faile with next command
//go run -race  .\racecondition\raicecondition.go
//-race allow detect race condition
//Found 3 data race(s)
