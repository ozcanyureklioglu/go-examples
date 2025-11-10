package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() string {
	return "Hav Benim adım " + d.Name
}

type Cat struct {
	Name string
}

func (c *Cat) Speak() string {
	return "Miyav Benim adım " + c.Name
}

func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func main(){
	dog := &Dog{Name: "Pamuk"}
	cat := &Cat{Name: "Cindy"}

	MakeSound(dog)
	MakeSound(cat)
}