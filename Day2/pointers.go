package main

import "fmt"

func pointerDemo() {
    x := 10
    fmt.Println("x before:", x)

    p := &x       // pointer to x
    *p = 20       // change x through pointer

    fmt.Println("x after:", x)

	name := "Cynthia"
	fmt.Println("Name before:", name)

	pName := &name	   // pointer to name
	*pName = "Mugo"     // change name through pointer
	fmt.Println("Name after:", name)
}
