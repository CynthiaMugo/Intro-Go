package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
	variables()
	conditionals()
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

func conditionals() {
    temperature := 15

    if temperature < 10 {
        fmt.Println("It's cold")
    } else if temperature < 20 {
        fmt.Println("Mild weather")
    } else {
        fmt.Println("Hot day")
    }
	// For loop example
	for i := 1; i <= 5; i++ {
    fmt.Println("Step", i)
	}

	// While loop example
	count := 1
	for count <= 5 {
		fmt.Println("Count:", count)
		count++
	}

	// switch case example
	fruit := "mango"

	switch fruit {
	case "apple":
		fmt.Println("Crisp and tasty")
	case "mango":
		fmt.Println("Sweet and juicy")
	default:
		fmt.Println("Unknown fruit")
	}

	// mini exercise
	score := 90

	if score < 50 {
		fmt.Println("Fail")
	} else if score < 75 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Distinction")
	}
	// Create a slice of numbers [1,2,3,4,5] and loop through them with a for loop, printing each.
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
	}  
	// corrected version by AI
	numbers := []int{1, 2, 3, 4, 5}
    for _, num := range numbers {
        fmt.Println("Number:", num)
    }

	// Make a day variable with a string and write a switch to print something for "Monday", "Friday", and a default for other days.
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the work week")
	default:
		fmt.Println("It's just another day")
	}
}
