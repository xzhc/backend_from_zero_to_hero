package practice

import (
	"fmt"
	"testing"
)

func produce(ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
}

func TestChannel(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: channel")

	ch := make(chan int)

	go func() {
		ch <- 23
	}()

	val := <-ch
	fmt.Println(val)

	ch1 := make(chan int, 10)

	ch1 <- 10
	ch1 <- 20
	ch1 <- 30
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	ch2 := make(chan int, 3)
	go produce(ch2)
	for val := range ch2 {
		fmt.Println("收到：", val)
	}
}
