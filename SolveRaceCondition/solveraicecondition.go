package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0

//wait group
var wg sync.WaitGroup

//mutex
var mu sync.Mutex

func goroutinefunc() {
	//Mutex to lock
	mu.Lock()
	v := counter
	//Yield the procesor
	runtime.Gosched()
	v++
	counter = v
	//Mutex to Unlock
	mu.Unlock()
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
