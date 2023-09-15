package main

import "fmt"

func main() {
	ch := make(chan int)
	signal := make(chan struct{})

	go send(ch)
	go receive(ch, signal)

	<-signal
}

func send(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i * i
	}
}

func receive(ch <-chan int, signal chan<- struct{}) {
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
	signal <- struct{}{}
}
