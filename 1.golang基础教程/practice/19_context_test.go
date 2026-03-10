package practice

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: context")

	// ctx, cancel := context.WithCancel(context.Background())

	// go func() {
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Println("goroutine收到取消信号，退出")
	// 			return
	// 		default:
	// 			fmt.Println("goroutine 运行中")
	// 			time.Sleep(500 * time.Millisecond)
	// 		}
	// 	}
	// }()

	// time.Sleep(1500 * time.Millisecond)
	// cancel()
	// time.Sleep(200 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("操作完成")
	case <-ctx.Done():
		fmt.Println("超时退出：", ctx.Err())
	}
}
