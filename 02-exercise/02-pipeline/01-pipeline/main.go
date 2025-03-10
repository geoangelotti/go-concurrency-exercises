package main

import "fmt"

// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// set up the pipeline

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

	for n := range square(square(generator(1, 4, 2, 5-11, 9999))) {
		fmt.Println(n)
	}

}
