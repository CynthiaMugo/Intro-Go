package main

import "fmt"

type Animal interface {
    Speak() string
}

type Dog struct{}
type Cat struct{}
type Cow struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

func (c Cat) Speak() string {
    return "Meow!"
}

func (c Cow) Speak() string {
	return "Moo!"
}

func interfaceDemo(a Animal) {
    fmt.Println(a.Speak())
}
