package main

import "fmt"

func main() {
	fmt.Println("Welcome to Day 2 of Go Programming!")
	maps()
	exercise()

	s := Student{
        Name:  "Cynthia",
        Age:   22,
        Grade: "A",
    }

    s.Introduce()

	c := Car {
		Brand: "Toyota",
		Year:  2020,
	}
	
	c.Display()
	c.UpdateYear(2022)
	c.Display()

	pointerDemo()

	interfaceDemo(Dog{})
	interfaceDemo(Cat{})
	interfaceDemo(Cow{})

}

func maps() {
	// Creating a map
	countryCapitalMap := make(map[string]string)

	// Adding key-value pairs to the map
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	// Retrieving a value from the map
	capital := countryCapitalMap["Japan"]
	fmt.Println("Capital of Japan is:", capital)

	// Deleting a key-value pair from the map
	delete(countryCapitalMap, "Italy")

	// Iterating over the map
	fmt.Println("Country Capitals:")
	for country, capital := range countryCapitalMap {
		fmt.Printf("%s: %s\n", country, capital)
	}

}

func exercise() {
    student := map[string]any{
		// any can hold any type of value it can be used instead of interface{}
        "name":  "Cynthia",
        "age":   22,
        "grade": "A",
    }

    fmt.Println(student)
	// student["school"] = "XYZ University"
    student["grade"] = "A+"

    val, exists := student["school"]
    fmt.Println("Exists?", exists, "Value:", val)

    delete(student, "grade")

    fmt.Println(student)
}