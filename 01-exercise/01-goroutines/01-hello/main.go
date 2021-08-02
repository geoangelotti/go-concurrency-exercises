package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// write goroutine with different variants for function call.

	// goroutine function call
	go fun("goroutine call")

	// goroutine with anonymous function
	go func(s string) {
		fun(s)
	}("anonymous call")

	// goroutine with function value call
	value := fun
	go value("function value call")

	// wait for goroutines to end
	time.Sleep(time.Second / 3)

	fmt.Println("done..")
}
