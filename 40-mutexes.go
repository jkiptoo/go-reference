package main

import (
	"fmt"
	"sync"
)

//Use mutexes to safely access data over multiple goroutines.
//Use mutex to manage complex states

//Use container struct to be updated from many goroutines. Since mutexes cannot be copied, use a pointer when passing the container struct.

type Container struct {
	containerMutex sync.Mutex
	counters map[string]int
}

//Lock containerMutex before accessing the counters.
func (c *Container) inc(name string) {
	//Lock containerMutex
	c.containerMutex.Lock()

	//Unlock containerMutex using defer statement
	defer c.containerMutex.Unlock()
	c.counters[name]++
}

func main() {
	//Zero value of mutex requires no initialisation
	c := Container {
		counters: map[string]int{"Alto": 0, "Beta": 0},
	}

	//Increment counter in loop
	var groupWait sync.WaitGroup

	increment := func(name string, nameCounter int) {
		for range nameCounter {
			c.inc(name)
		}
		groupWait.Done()
	}

	//Concurrently run goroutines together with all of them accessing the same container and two accessing the same counter
	groupWait.Add(3)
	go increment("Alto", 12000)
	go increment("Alto", 12000)
	go increment("Beta", 12500)

	//Wait for goroutines to complete
	groupWait.Wait()
	fmt.Println(c.counters)
}