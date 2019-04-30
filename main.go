package main

import (
	"fmt"
)

const (
	// Number seats at the table
	seats  = 5
	// Number of portions on each plate
	portions = 3
)

func main() {
	// Create a waiter
	waiter := &Waiter{
		// Number of philosophers this waiter can serve at once
		serves: make(chan bool),
		// Number of requests a waiter can take
		requests: make(chan bool),
		// Flag for service finished
		done: make(chan bool),
	}

	// Create seats * chopsticks
	chopSticks := make([]*ChopStick, seats)
	for i := 0; i < seats; i++ {
		chopSticks[i] = new(ChopStick)
	}

	// Create seats * philosophers
	philosophers := make([]*Philosopher, seats)
	for i := 0; i < seats; i++ {
		philosophers[i] = &Philosopher{
			id:     i + 1,
			left:   chopSticks[i],
			right:  chopSticks[(i+1)%seats],
			waiter: waiter,
			food:   Food{portions},
			done:   make(chan bool),
		}
	}

	fmt.Println("Table Service Has Started!")

	// Start serving the philosophers
	waiter.serve(philosophers)

	fmt.Println("Table Service Has Finished!")
}
