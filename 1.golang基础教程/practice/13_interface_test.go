package practice

import (
	"fmt"
	"testing"
)

type Speaker interface {
	Sound() string
}

type IntDog struct{}

type IntCat struct{}

func (d IntDog) Sound() string {
	return "汪汪！"
}

func (c IntCat) Sound() string {
	return "喵喵！"
}

func makeSound(a Speaker) {
	fmt.Println(a.Sound())
}

func TestInterface(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: interface")

	dog := IntDog{}
	cat := IntCat{}

	makeSound(dog)
	makeSound(cat)
}
