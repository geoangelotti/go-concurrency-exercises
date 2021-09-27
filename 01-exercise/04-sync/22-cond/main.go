package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		cond.L.Lock()
		for len(sharedRsc) < 2 {
			cond.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		cond.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		cond.L.Lock()
		for len(sharedRsc) < 2 {
			cond.Wait()
		}

		fmt.Println(sharedRsc["rsc2"])
		cond.L.Unlock()
	}()

	// writes changes to sharedRsc
	go func() {
		cond.L.Lock()
		sharedRsc["rsc1"] = "foo"
		cond.Broadcast()
		cond.L.Unlock()
	}()

	go func() {
		cond.L.Lock()
		sharedRsc["rsc2"] = "bar"
		cond.Broadcast()
		cond.L.Unlock()
	}()

	wg.Wait()
}
