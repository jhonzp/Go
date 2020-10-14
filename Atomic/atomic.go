package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64 = 0

//wait group
var wg sync.WaitGroup

func goroutinefunc() {

	v := atomic.LoadInt64(&counter)
	//v:=counter
	//Yield the procesor
	runtime.Gosched()
	v++
	//set value
	//counter = v
	atomic.StoreInt64(&counter, v)

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
