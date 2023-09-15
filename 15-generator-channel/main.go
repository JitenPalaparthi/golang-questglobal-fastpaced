package main

import "fmt"

func main() {
	sig1 := make(chan struct{})
	sig2 := make(chan struct{})

	ch1 := send(20, "sender-1")
	ch2 := send(20, "sender-2")

	go receive(ch1, sig1, "receiver-1")
	go receive(ch2, sig2, "receiver-2")

	<-sig1
	<-sig2
}

func send(num int, sender string) chan string {
	ch := make(chan string)

	go func() {
		for i := 1; i <= num; i++ {
			ch <- fmt.Sprint(sender, "-->", i*i)
		}
		close(ch)
	}()

	return ch
}

func receive(ch <-chan string, signal chan<- struct{}, routine string) {
	for sq := range ch {
		fmt.Println("received by ", routine, sq)
	}
	signal <- struct{}{}
}
