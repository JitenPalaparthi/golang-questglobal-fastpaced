package main

import "fmt"

func main() {
	ch := make(chan int)
	signal := make(chan struct{})

	go send(ch, 40)

	go receive1(ch, "receiver-1", signal)
	go receive2(ch, "receiver-2", signal)

	<-signal
}

func send(ch chan int, num int) {
	for i := 1; i <= num; i++ {
		ch <- i * i
	}
	close(ch)
}

func receive1(ch <-chan int, routine string, signal chan<- struct{}) {
	for {
		sq, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("received by ", routine, sq)
	}
	signal <- struct{}{}
}

func receive2(ch <-chan int, routine string, signal chan<- struct{}) {
	for sq := range ch {
		fmt.Println("received by ", routine, sq)
	}
	signal <- struct{}{}
}
