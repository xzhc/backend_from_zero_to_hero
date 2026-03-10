package practice

import (
	"fmt"
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: mutex")

	var mu sync.Mutex
	var count int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	t.Log("Count = ", count)
}
