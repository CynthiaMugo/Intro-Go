package main

import "fmt"

type Car struct {
    Brand string
    Year  int
}

// method with value receiver
func (c Car) Display() {
    fmt.Printf("Brand: %s, Year: %d\n", c.Brand, c.Year)
}

// method with pointer receiver (can modify)
func (c *Car) UpdateYear(newYear int) {
    c.Year = newYear
}
