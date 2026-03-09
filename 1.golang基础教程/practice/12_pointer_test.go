package practice

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: pointer")

	i := 23

	p := &i
	fmt.Println("i的地址是：", p)
	fmt.Println("i的值是：", *p)

	//*ptr
	*p = 100
	fmt.Println("修改后i的值是：", i)
}
