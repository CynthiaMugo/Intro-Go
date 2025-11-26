# Go Language Learning – Day 1 & Day 2

A simple collection of Go exercises and examples documenting my learning journey.  
This project includes variables, conditionals, functions, loops, arrays, slices, maps, structs, methods, interfaces, pointers, and multi-file organization.

---

## Project Overview

This repository contains practice code for learning the basics of the Go programming language.  
All exercises are structured into daily folders (e.g., `Day1`, `Day2`) each containing a `main.go` file and additional Go files as needed.

The project is intentionally simple and beginner-friendly, focusing on understanding Go syntax and core concepts.

---

## Folder Structure

```
Go-language/
│
├── Day1/
│   ├── main.go
│   ├── functions.go
│   └── (other concept files)
│
└── Day2/
    ├── main.go
    ├── structs.go
    ├── methods.go
    └── (more files as concepts grow)
```

## Day 1 – Fundamentals
### Variables & Constants

Declaring variables with :=
Explicit type declarations
Working with constants
Printing types using fmt.Printf("%T")

### Conditionals & Loops

if, else if, else
for loops
Range loops (for _, val := range slice)
switch statements
Mini exercises like score grading

### Functions

Simple functions
Functions with parameters
Return values
Multiple return values (with error handling!)
Calling functions from another file in the same package

### Multi-File Execution

Using:
```
go run .
```

## Day 2 – Structs, Methods & Interfaces
### Structs
Defining custom data types using struct.

###  Methods
Attaching functions to structs using value or pointer receivers.

### Pointers
Understanding how Go handles memory references and modifying actual values.

### Interfaces
Using Go's implicit interface system — no need for implements.

### Organizing Code
Splitting structs, methods, and logic into multiple files within the same folder.

## Running the Project
Navigate into any day folder, for example:
```
cd Day1
go run .
```

Go will automatically execute all .go files in that folder (as long as they use package main).

## Purpose

This project is part of my hands-on journey to master Go — learning through building small examples, experimenting, observing outputs, and gradually moving toward real-world Go applications.