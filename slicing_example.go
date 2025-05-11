package main

import "fmt"

func main() {
	// Create a slice
	nums := []int{1, 2, 3, 4, 5}

	// Access elements
	fmt.Println(nums[1]) // Output: 2

	// Slicing
	sub := nums[1:4] // Slice from index 1 to 3 (excludes 4)
	fmt.Println(sub) // Output: [2 3 4]

	// Append to slices
	nums = append(nums, 6, 7)
	fmt.Println(nums)      // Output: [1 2 3 4 5 6 7]
	fmt.Println(cap(nums)) // Output: [1 2 3 4 5 6 7]
}
