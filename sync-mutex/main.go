package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()

	lock.Lock()
	balance += amount
	lock.Unlock()
}

func PrintBalance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()

	return b
}

func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}

	wg.Wait()

	fmt.Println("Balance:", PrintBalance(&lock))
}
