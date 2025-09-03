package main

import "sync"

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

	//Increment counter in loops
	var groupWait sync.

}