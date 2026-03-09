package practice

import (
	"fmt"
	"testing"
)

type Engine struct {
	Horsepower int
}

func (e Engine) Start() {
	fmt.Println("引擎启动")
}

type Car struct {
	Engine
	Brand string
}

func TestEmbed(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: embed")

	c := Car{
		Engine: Engine{Horsepower: 1000},
		Brand:  "BMW",
	}

	fmt.Println(c.Horsepower)
	c.Start()
}
