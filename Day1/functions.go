package main

import "fmt"

// mini-exercise
// A function introduce(name string, age int) that prints:
// "My name is X and I am Y years old."

// A function multiply(x, y int) int
// and call it inside basicsFunctions().

// A function isEven(num int) bool
// and print true/false.

// Simple function with no parameters
func greet() {
    fmt.Println("Hello from a function!")
}

// Function with parameters
func add(a int, b int) int {
    return a + b
}

func introduce(name string, age int) {
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)
}

func multiply(x int, y int) int {
	return x * y
}

func isEven(num int) bool {
	return num%2 == 0
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func basicsFunctions() {
    fmt.Println("SECTION 4: FUNCTIONS")

    greet()
	introduce("Lewis", 34)

	multiple := multiply(4, 5)
	fmt.Println("Multiplication Result:", multiple)

	boole := isEven(100)
	boolean := isEven(77)
	fmt.Println("Is 100 even:", boole, "and is 77 even:", boolean)

    sum := add(10, 20)
    fmt.Println("Sum:", sum)

    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Division Result:", result)
    }
}

