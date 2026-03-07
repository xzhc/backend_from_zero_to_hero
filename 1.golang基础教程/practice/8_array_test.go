package practice

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: array")

	var a [5]int
	a[4] = 100
	fmt.Println(a)

	// b := [5]int{1, 2, 3, 4, 5}
	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := [...]int{1, 2, 3}
	fmt.Println(c)
	fmt.Println(len(c))
}
