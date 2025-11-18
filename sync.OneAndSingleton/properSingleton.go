package main

import (
	"fmt"
	"sync"
)

func properSingleton() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := GetProperInstance()
			fmt.Printf("Singleton instance address: %p\n", s)
		}()
	}

	wg.Wait()
}
