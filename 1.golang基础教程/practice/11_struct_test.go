package practice

import (
	"fmt"
	"testing"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func (u User) Greet() string {
	return "Hello, my name is " + u.Name
}

func (u *User) Birthday() {
	u.Age += 1
}

func TestStruct(t *testing.T) {
	// TODO: 在这里写你的练习代码
	fmt.Println("practice: struct")

	u1 := User{
		Name:  "xzh",
		Age:   22,
		Email: "xzh@example.com",
	}
	fmt.Println(u1)

	u2 := User{"xzh", 22, "xzh@example.com"}
	fmt.Println(u2)

	u1.Age = 26
	fmt.Println(u1)

	u3 := u1
	u3.Name = "xzh2"
	fmt.Println(u1)
	fmt.Println(u3)

	//if you want to share the same memory, use pointer
	u4 := &u1
	u4.Name = "xzh3"
	fmt.Println(u1)
	fmt.Println(u4)

	//bind a method to struct
	//value receiver: Greet() defined at package level
	//pointer receiver: Birthday() defined at package level

	fmt.Println(u1.Greet())
	u1.Birthday()
	fmt.Println(u1)
}
