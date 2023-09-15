package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

func main() {

	wg := new(sync.WaitGroup)

	//wg.Add(3)

	go func(*sync.WaitGroup) {
		wg.Add(1)
		sq(100)
		wg.Done() // done is counter-1
	}(wg)
	// t1

	go func(*sync.WaitGroup) {
		wg.Add(1)
		sq(101)
		wg.Done()
	}(wg)
	// t1

	go func(*sync.WaitGroup) {
		wg.Add(1)
		sq(102)
		wg.Done()
	}(wg)
	// t1

	//go block() // t1
	// go sq(103) // t1
	// go block()
	// go block()
	// go block()
	// go sq(100) // ?
	fmt.Println("Main has ended")

	//runtime.Goexit()

	//time.Sleep(time.Millisecond * 100)

	wg.Wait() // wait until the wg counter is zero
}

func sq(num int) {
	fmt.Println(num * num)
}

func block() {
	counter := 1
	for {
		fmt.Println(counter)
		if counter > 10 {
			runtime.Goexit()
		}
		counter++
	}
}

func fatal(msg string) {
	fmt.Println(msg)
	os.Exit(1) // it exists the application as a failure
}

// err
// panic
// fatal

// go
// main is also a go routine
// no goroutine waits for other goroutine(s) to complete its execution
// go routines are executed on threads
// number of threads are created by go runtime is equal to num of CPU
