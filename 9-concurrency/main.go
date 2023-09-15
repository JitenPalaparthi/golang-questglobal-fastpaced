package main

import (
	"fmt"
	"time"
)

// Do not communicate by sharing memory; instead, share memory by communicating.

func main() {
	var ch chan int // 8 bytes o 64 bit machine
	ch = make(chan int)
	go func() {
		time.Sleep(time.Second * 1)
		ch <- 10
	}()
	fmt.Println(time.Now())
	signal := make(chan struct{})

	go func() {
		b := <-ch
		println(b)
		signal <- struct{}{}
		//return signal
	}()

	fmt.Println(time.Now())
	<-signal
}

// to establish comminication between multiple go routines channels are used
// chan is a keyword used to declare a channel
// make is builtin function to instantiate a channel/
// a channel can be buffered or unbuffered.
// most of the cases we use unbuffered channels. There are special or requirement specific cases to use buffered channels
// sending a value to the channel   ch <-10
// receiving a value from a channel  <-ch
// Note: For unbuffered channels
//		The sender is blocked until the receiver receives the value
// 		The receiver is blocked until the sender sends the value
// emtpty struct channel is used for signal or notification but it does not carry any data
