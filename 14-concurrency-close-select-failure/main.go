package main

import "fmt"

func main() {
	ch := make(chan int)

	signal := make(chan struct{})
	closeSig := make(chan struct{})

	go send(ch, closeSig, 40)

	go receive2(ch, signal, closeSig, "receiver-1")
	go receive2(ch, signal, closeSig, "receiver-2")

	<-signal
}

func send(ch chan int, closeSig <-chan struct{}, num int) {
	for i := 1; i <= num; i++ {
		select {
		case ch <- i * i:
		case <-closeSig:
			break
		}
	}
	close(ch)
}

func receive1(ch <-chan int, signal chan<- struct{}, closeSig chan<- struct{}, routine string) {
	for {
		sq, ok := <-ch
		if !ok {
			close(closeSig)
			break
		}
		fmt.Println("received by ", routine, sq)
	}
	signal <- struct{}{}
}

func receive2(ch <-chan int, signal chan<- struct{}, closeSig chan struct{}, routine string) {
	for sq := range ch {
		fmt.Println("received by ", routine, sq)
	}
	// select {
	// case _, ok := <-closeSig:
	// 	if ok {
	// 		close(closeSig)
	// 	}
	// }
	close(closeSig)
	signal <- struct{}{}
}
