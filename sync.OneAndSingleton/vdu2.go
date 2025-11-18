package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	counter int
}

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		fmt.Printf("...get instance...\n")
		instance = &Singleton{}
	}

	return instance
}

func vdu2() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := GetInstance()
			fmt.Printf("Singleton instance address: %p\n", s)
		}()
	}

	wg.Wait()
}
