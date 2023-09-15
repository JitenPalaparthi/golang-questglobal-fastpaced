package main

import (
	"fmt"
	"runtime"
)

func main() {
	//sig1 := make(chan struct{})
	//sig2 := make(chan struct{})

	ch1 := send(20, "sender-1")
	ch2 := send(20, "sender-2")
	// ch3 := send(20, "sender-3")

	// go receive(ch1, sig1, "receiver-1")
	// go receive(ch2, sig2, "receiver-2")
	// <-sig1
	// <-sig2
	// <-receive2(ch3, "receiver-3") // This is not a async operation
	sig1, sig2 := receive3(ch1, ch2)
	// sig1, sig2 := receive4(ch1, ch2)
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
	signal <- struct{}{} // instead of it can also use close(singal)
}

func receive2(ch <-chan string, routine string) chan struct{} {
	signal := make(chan struct{}, 1)
	// for sq := range ch {
	// 	fmt.Println("received by ", routine, sq)
	// }

	go func() {
		defer fmt.Println("This go routine is exiting")
		//outer:
		for {
			select {
			case sq, ok := <-ch:
				if !ok {
					close(signal) // close channel also sends a signal
					runtime.Goexit()
					//break outer
				} else {
					fmt.Println("received by ", routine, sq)
				}
			}
		}
	}()
	return signal
}

func receive3(ch1, ch2 chan string) (sig1 chan struct{}, sig2 chan struct{}) {
	sig1 = receive2(ch1, "receiver1 ch1")
	sig2 = receive2(ch2, "receiver1 ch1")
	return sig1, sig2
}

// func receive4(ch1, ch2 chan string) (sig1 chan struct{}, sig2 chan struct{}) {
// 	signal1 := make(chan struct{}, 1)
// 	signal2 := make(chan struct{}, 1)
// 	go func() {
// 		counter := 0
// 		for {
// 			select {
// 			case sq, ok := <-ch1:
// 				if !ok {
// 					counter++
// 					//close(signal1) // close channel also sends a signal
// 					//runtime.Goexit()
// 					//break outer
// 					//signal1 <- struct{}{}
// 				} else {
// 					fmt.Println("received by ch1->", sq)
// 				}

// 			case sq, ok := <-ch2:
// 				if !ok {
// 					counter++
// 					//close(signal2) // close channel also sends a signal
// 					//runtime.Goexit()
// 					//break outer
// 					//signal2 <- struct{}{}
// 				} else {
// 					fmt.Println("received by ch2->", sq)
// 				}
// 			default:
// 				if counter == 2 {
// 					runtime.Goexit()
// 				}
// 			}
// 		}
// 	}()

// 	return signal1, signal2

// }
