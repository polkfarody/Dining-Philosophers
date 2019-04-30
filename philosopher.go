package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	eat   = 150 // Time spent eating in milliseconds
	think = 100 // Time spent thinking in milliseconds
)

// A philosopher has an ID a chopstick for each hand,
// access to the waiter and a plate of food
type Philosopher struct {
	id          int
	left, right *ChopStick
	waiter      *Waiter
	done        chan bool
	food        Food
}

// This philosopher will have a session of thinking and eating until
// all food portions are gone.
func (p *Philosopher) start() {
	for ; p.food.portions > 0; p.food.portions-- {
		p.timeout(think)
		p.eat()
	}

	// Send message into channel to acknowledge this philosopher has finished
	p.done <- true
}

// Before a philosopher can eat it needs permission from the waiter.
// Once granted permission the philosopher can pick up two chopsticks
// and begin to eat.
func (p *Philosopher) eat() {
	p.waiter.askPermissionToEat(p)
	fmt.Println("Phil", p.id, "is eating...")
	// Take some time eating.
	p.timeout(eat)

	fmt.Println("Phil", p.id, "finished")
	p.waiter.onFinishEating(p)
}

// Sets a timeout for a specific duration
func (p *Philosopher) timeout(duration time.Duration) {
	time.Sleep(duration * time.Millisecond)
}

// Put down chopsticks in a random order
func (p *Philosopher) putDownChopSticks() {
	switch rand.Intn(2) {
	case 0:
		p.left.Unlock()
		p.right.Unlock()
	case 1:
		p.right.Unlock()
		p.left.Unlock()
	}
}

// Pick up chopsticks in a random order
func (p *Philosopher) pickUpChopSticks() {
	switch rand.Intn(2) {
	case 0:
		p.left.Lock()
		p.right.Lock()
	case 1:
		p.right.Lock()
		p.left.Lock()
	}
}
