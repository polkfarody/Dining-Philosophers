package main

import "sync"

// The role of a chopstick is to lock when picked up and
// unlock when put back down
type ChopStick struct {
	sync.Mutex
}

// An amount of portions
type Food struct {
	portions int
}