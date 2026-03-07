package practice

import (
	"fmt"
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: switch")

	var i = 2
	switch i {
	case 1:
		fmt.Println("i的值为1")

	case 2:
		fmt.Println("i的值为2")

	case 3:
		fmt.Println("i的值为3")
	}

	time := time.Now()

	switch {
	case time.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
