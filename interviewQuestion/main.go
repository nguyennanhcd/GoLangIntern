package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var mu sync.Mutex

var counter int64

// dùng mutex
func mutexIncrement() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

// dùng atomic
func atomicIncrement() {
	atomic.AddInt64(&counter, 1)
}

func main() {
	var wg sync.WaitGroup

	// ở đây chỉ tạo 1000 goroutine chứ không chạy liền
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomicIncrement()
		}()
	}

	// bắt main go routine chờ tất cả các goroutine hoàn thành
	wg.Wait()
	fmt.Println("Final counter:", counter)
}
