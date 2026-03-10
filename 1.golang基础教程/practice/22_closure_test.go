package practice

import (
	"fmt"
	"testing"
)

func outer() func() int {

	x := 0

	return func() int {
		x++
		return x
	}
}

func TestClosure(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: closure")
	f := outer()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
