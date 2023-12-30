package main

import (
	"fmt"
)

// This recursive algorithm prints the integer n if n > 1, then calls itself with n / 2 as the argument until n <= 1.
// Input: n int
// Output: None (prints the integer n)
func recursiveAlgorithm(n int) {
	// Base case
	if (n <= 1) {
		return;
	}
	// Recursive case
	recursiveAlgorithm(n / 2);
	// Print the integer n
	fmt.Println(n);
}

func main() {
	// Test case 0
	fmt.Println("Test case 0:");
	n := 9;
	recursiveAlgorithm(n); 

	fmt.Println("----------------------------------------");

	// Test case 1
	fmt.Println("Test case 1:");
	n = 16;
	recursiveAlgorithm(n);

	fmt.Println("----------------------------------------");

	// Test case 2
	fmt.Println("Test case 2:");
	n = 77;
	recursiveAlgorithm(n);

	fmt.Println("----------------------------------------");

	// Test case 3
	fmt.Println("Test case 3:");
	n = 73;
	recursiveAlgorithm(n);

	fmt.Println("----------------------------------------");

	// Test case 4
	fmt.Println("Test case 4:");
	n = 1;
	recursiveAlgorithm(n);

}