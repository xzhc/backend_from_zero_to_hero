package practice

import (
	"fmt"
	"testing"
)

// 类型和方法必须定义在包级别，不能嵌套在函数内部
type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "wang wang"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "miao miao"
}

type Celsius float64
type Fahrenheit float64

type MyString string

func TestType(t *testing.T) {
	// TODO: 在这里写你的练习代码
	type Person struct {
		name string
		age  int
	}

	p := Person{
		name: "Alice",
		age:  25,
	}
	fmt.Println(p.name, p.age)

	animals := []Animal{Dog{}, Cat{}}
	for _, a := range animals {
		fmt.Println(a.Speak())
	}

	var c Celsius = 100.0
	var f Fahrenheit = 212.0

	f = Fahrenheit(c)
	fmt.Println(f)

	var s MyString = "hello"
	fmt.Println(s)

}
