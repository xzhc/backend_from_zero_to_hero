package practice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: slice")

	var s1 []string
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := make([]string, 3, 5)
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	s3 := []string{"a", "b", "c"}
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))

	s := []string{"a", "b", "c"}
	s = append(s, "d")
	s = append(s, "e", "f")
	s4 := []string{"g", "h"}
	s = append(s, s4...)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s6 := []int{1, 2, 3, 4, 5}
	s7 := s6[1:3]
	s6[0] = 99
	fmt.Println(s6)
	fmt.Println(s7)
	fmt.Println(len(s7))
	fmt.Println(cap(s7))

	s8 := make([]int, 2)
	copy(s8, s6[1:3])
	fmt.Println(s8)
	fmt.Println(len(s8))
	fmt.Println(cap(s8))
}
