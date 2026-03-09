package practice

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func say(msg string) {
	fmt.Println(msg)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d 开始工作\n", id)
	fmt.Printf("Worker %d 完成工作\n", id)
}

func TestGoroutine(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: goroutine")

	go say("在新的goroutine里执行")
	say("在 主goroutine里执行")

	time.Sleep(time.Millisecond * 100)

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("所有 worker 都完成了")
}
