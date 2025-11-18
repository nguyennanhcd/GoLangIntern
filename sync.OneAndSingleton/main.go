package main

import (
	"fmt"
	"sync"
)

type ProperSingleton struct{}

var (
	properInstance *ProperSingleton
	properOnce     sync.Once
)

func GetProperInstance() *ProperSingleton {
	properOnce.Do(func() {
		fmt.Printf("...get proper instance...\n")
		properInstance = &ProperSingleton{}
	})
	return properInstance
}
func main() {
	// vdu1()
	// vdu2()
	// syncOne()
	properSingleton()
}
