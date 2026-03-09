package practice

import (
	"fmt"
	"testing"
	"time"
)

func slowOperation(ch chan string) {
	time.Sleep(time.Second * 5)

	ch <- "操作结果"
}

func TestSelect(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: select")

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)

		ch1 <- "goroutine1 完成"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "goroutine2 完成"
	}()

	// ch1 <- "来自ch1的消息"
	// ch2 <- "来自ch2的消息"

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("收到：", msg)

		case msg := <-ch2:
			fmt.Println("收到：", msg)
		}
	}

	ch3 := make(chan string)

	select {
	case result := <-ch3:
		fmt.Println("成功：", result)
	case <-time.After(time.Second * 3):
		fmt.Println("超时，等待超过了3s")
	}

}
