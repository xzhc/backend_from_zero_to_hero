package practice

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: for")

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 1
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	for i := range 3 {
		fmt.Println("range:", i)
	}

	// for {
	// 	fmt.Println("死循环")
	// }
}
