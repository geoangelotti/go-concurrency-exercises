package main

import "fmt"

// Implement relaying of message with Channel Direction

func genMsg(ch chan<- string) {
	// send message on ch1
	ch <- "hello"
}

func relayMsg(in <-chan string, out chan<- string) {
	// recv message on ch1
	r := <-in
	// send it on ch2
	out <- r
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spine goroutine genMsg and relayMsg
	go genMsg(ch1)
	go relayMsg(ch1, ch2)
	// recv message on ch2
	r := <-ch2
	fmt.Println(r)
}
