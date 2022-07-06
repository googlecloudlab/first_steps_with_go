package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var balance int64

// To implement mutual exclusion, Go's standard library provides the Mutex type
// in the sync package. A Mutex is a mutual exclusion lock. A mutual exclusion
// lock is a technique to ensure exclusive access to shared data between threads
// of execution. When a Goroutine gains the mutual exclusion lock, other Goroutines
// need to wait until the lock is released.
var mutex = &sync.Mutex{}

func credit(wg *sync.WaitGroup) {
	// notify the WaitGroup when we are done
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()

		// adds 100 to balance atomically
		// atomic operations allow to perform mathematical operations on variables one thread at a time
		atomic.AddInt64(&balance, 100)

		fmt.Println("After crediting, balance is", balance)
		mutex.Unlock()

		time.Sleep((time.Duration(rand.Intn(100)) * time.Millisecond))
	}
}

func debit(wg *sync.WaitGroup) {
	// notify the WaitGroup when we are done
	defer wg.Done()
	for i := 0; i < 5; i++ {
		mutex.Lock()

		// deducts -100 from balance atomically
		// atomic operations allow to perform mathematical operations on variables one thread at a time
		atomic.AddInt64(&balance, -100)

		fmt.Println("After debiting, balance is", balance)
		mutex.Unlock()

		time.Sleep((time.Duration(rand.Intn(100)) * time.Millisecond))
	}
}

func main() {
	// Create a WaitGroup object
	var wg sync.WaitGroup

	balance = 200
	fmt.Println("Initial balance is", balance)

	// Calling the wg.Add() function to increment the WaitGroup counter prior to calling the credit() goroutine
	wg.Add(1) // add 1 to the WaitGroup counter
	go credit(&wg)

	// Calling the wg.Add() function to increment the WaitGroup counter prior to calling the debit() goroutine
	wg.Add(1) // add 1 to the WaitGroup counter
	go debit(&wg)

	// Wait for the completion of the two goroutines
	wg.Wait() // blocks until WaitGroup counter is 0

	fmt.Println("Ending balance is", balance)
}
