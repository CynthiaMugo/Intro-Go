package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
	variables()
}
	

func variables() {
	// Variable declarations
	age := 25
	// explicit type declaration
	var age1 int = 27
	name := "Cynthia"
	isSoftwareEngineer := true
	height := 5.4
	pi := 3.142
	// we can also use const for constant values
	const piValue = 3.14159
	
	// Print the constant value
	fmt.Println("Pi Constant:", piValue)

	// Print the variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Age1:", age1)
	fmt.Println("Is Software Engineer:", isSoftwareEngineer)
	fmt.Println("Height:", height)
	fmt.Println("Pi:", pi)
	// Print the types of the variables
	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of isSoftwareEngineer: %T\n", isSoftwareEngineer)
	fmt.Printf("Type of height: %T\n", height)
	fmt.Printf("Type of pi: %T\n", pi)
}
