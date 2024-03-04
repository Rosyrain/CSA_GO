package main

import (
	"fmt"
	"sync"
)

var (
	count int32
	wg    sync.WaitGroup
	Mutex sync.Mutex
)

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			Mutex.Lock()
			defer Mutex.Unlock()
			defer wg.Done()
			fmt.Println(count, "你好")
			count += 1
		}()

	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			Mutex.Lock()
			defer Mutex.Unlock()
			defer wg.Done()
			fmt.Println(count, "我不好")
			count += 1
		}()

	}
	wg.Wait()
	return
}
