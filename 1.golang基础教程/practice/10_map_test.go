package practice

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: map")

	//1. use make to create a map(recommend, declare first, then fill in data)
	scores1 := make(map[string]int)

	scores1["Alice"] = 95
	scores1["Bob"] = 87
	scores1["Bob"] = 88
	fmt.Println(scores1)

	//2. use map literal to create a map
	scores2 := map[string]int{
		"Alice": 95,
		"Bob":   87,
	}
	fmt.Println(scores2)

	//3.Query
	//3.1 get the value directly(if the key isn't exist, it will return the zero value. It's 0 for int)
	v := scores1["Alice"]
	fmt.Println(v)

	//3.2 use two-value form to query
	v1, ok := scores2["xzh"]
	if ok {
		fmt.Println("v1 is:", v1)
	} else {
		fmt.Println("Bob isn't exist")
	}

	//4. delete
	delete(scores1, "Bob")
	fmt.Println(scores1)

	//5.clear
	clear(scores2)
	fmt.Println(scores2)
}
