package main

import (
	"fmt"
	"time"
)

func vdu1() {
	// vdu1:

	s := GetInstance()
	fmt.Printf("Singleton instance address: %p\n", s)

	time.Sleep(2 * time.Second)

	f := GetInstance()
	fmt.Printf("Singleton instance address: %p\n", f)
}
