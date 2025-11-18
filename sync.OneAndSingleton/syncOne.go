package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("...Initializing...")
}

func syncOne() {
	for i := 0; i < 5; i++ {
		once.Do(initialize)
		fmt.Println("-->", i+1, "-->")
	}
}
