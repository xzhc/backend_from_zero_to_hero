package practice

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: ticker")

	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	fmt.Println("等待2s...")
	<-timer.C
	fmt.Println("2s到了")

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for i := 0; i < 3; i++ {
		<-ticker.C
		fmt.Printf("第%d次触发\n", i+1)
	}
}
