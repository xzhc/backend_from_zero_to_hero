package practice

import (
	"fmt"
	"testing"
)

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func logDecorator(fn func(string)) func(string) {
	return func(name string) {
		fmt.Println("[装饰器] 调用前")
		fn(name)
		fmt.Println("[装饰器] 调用后")
	}
}

func TestDecorator(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: decorator")

	decoratot := logDecorator(sayHello)
	decoratot("Alice")
}
