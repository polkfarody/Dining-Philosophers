package main

type Waiter struct {
	requests chan bool
	serves   chan bool
	done     chan bool
}

// Starts the table service by starting all philosophers
// It then waits for a request of service and then continues to serve
// Once all of the philosophers have eaten the service is complete and returns
func (w *Waiter) serve(philosophers []*Philosopher) {
	w.start(philosophers)

	for {
		select {
		case <-w.done:
			return

		case <-w.requests:
			w.serves <- true
		}

	}
}

// Forces each philosopher to eat and then waits for them to finish
func (w *Waiter) start(philosophers []*Philosopher) {
	for _, p := range philosophers {
		go p.start()
	}

	go w.wait(philosophers)
}

// Waits for each philosopher to finish and then sends the flag to the waiter.
func (w *Waiter) wait(philosophers []*Philosopher) {
	for _, p := range philosophers {
		<-p.done
	}

	w.done <- true
}

// A philosopher must send a request to the waiter and then wait for service.
// Once that is complete they can pick up their chopsticks
func (w *Waiter) askPermissionToEat(p *Philosopher) {
	w.requests <- true
	<-w.serves
	p.pickUpChopSticks()

}

// Put down the chopsticks
func (w *Waiter) onFinishEating(p *Philosopher) {
	p.putDownChopSticks()
}
