package main

import "fmt"

// Mini Exercise (your turn)
// Add inside arraysAndSlices():
// Create an array of 4 city names
// Create a slice with your top 2 favorite cities
// Append 1 more city
// Print all three with labels

func arraysAndSlices() {
    fmt.Println("SECTION 5: ARRAYS & SLICES")

    // --- Arrays ---
    var fixedArr [3]int = [3]int{10, 20, 30}
    fmt.Println("Array:", fixedArr)

    // --- Slices ---
    nums := []int{1, 2, 3}
    fmt.Println("Slice:", nums)

    // Append (adds elements)
    nums = append(nums, 4, 5)
    fmt.Println("After append:", nums)

    // Slice a slice
    part := nums[1:4]
    fmt.Println("Sliced section:", part)

    // Length and capacity
    fmt.Println("Length:", len(nums))
    fmt.Println("Capacity:", cap(nums))
}
