package practice

import (
	"fmt"
	"testing"
)

func greet(name string) string {
	return "Hello, " + name
}

func add(a, b int) int {
	return a + b
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func TestFunc(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println(greet("xzh"))
	fmt.Println(add(1, 2))
	quotient, remainder := divide(10, 3)
	fmt.Println(quotient, remainder)

	fmt.Println(minMax([]int{1, 2, 3, 4, 5}))

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))

	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))

	result := func(x int) int {
		return x * x
	}(5)
	fmt.Println(result)

	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

}
