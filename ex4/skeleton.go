package main

import "sync"

type MutexCounter struct {
	lck   sync.Mutex
	count int
}

func (c *MutexCounter) Add() {
	// Your code here
}

func (c *MutexCounter) Get() int {
	// Your code here
	return 0
}

type RWMutesCounter struct {
	lck   sync.RWMutex
	count int
}

func (c *RWMutesCounter) Add() {
	// Your code here
}

func (c *RWMutesCounter) Get() int {
	// Your code here
	return 0
}
