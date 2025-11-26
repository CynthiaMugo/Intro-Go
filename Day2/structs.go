package main

import "fmt"

type Student struct {
    Name  string
    Age   int
    Grade string
}

// Value receiver — gets a COPY of the struct
func (s Student) Introduce() {
    fmt.Printf("Hi, I'm %s, %d years old and in grade %s.\n", s.Name, s.Age, s.Grade)
}
