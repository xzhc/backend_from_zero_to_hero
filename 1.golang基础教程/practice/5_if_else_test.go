package practice

import (
	"fmt"
	"testing"
)

func TestIf_else(t *testing.T) {
	// TODO: 在这里写你的练习代码
	if 7%2 == 0 {
		fmt.Println("7 是偶数")
	} else {
		fmt.Println("7 是奇数")
	}
}
