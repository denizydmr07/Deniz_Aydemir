package main

import (
	"fmt"
)

// This function takes one parameter as an array/list. Find most repeated data within a given array.
// If there are multiple data with the same number of repetitions, returns the first one.
// If all data are unique, returns the first one.
// Input: words []string, 
// Output: string, most repeated word
func mostRepeatedWord(words []string) string {
	// Declare a map to store the words and their counts
	wordCount := make(map[string]int)

	// Iterate through the words, ignore the index
	for _, word := range words {
		// If the word is already in the map, increment its count
		// Using the ok variable to check if the word is in the map
		if _, ok := wordCount[word]; ok {
			wordCount[word]++
		} else {
			// If the word is not in the map, add it to the map with a count of 1
			wordCount[word] = 1
		}
	}

	// Declare a variable to store the most repeated word, initialize it to the first word in the array
	mostRepeatedWord := words[0]

	// Iterate through the map, find the most repeated word
	for word, count := range wordCount {
		// If the count of the current word is greater than the count of the most repeated word, update the most repeated word
		if count > wordCount[mostRepeatedWord] {
			mostRepeatedWord = word
		}
	}

	// Return the most repeated word
	return mostRepeatedWord
}

func main() {
	// Test case 0
	words := []string{"apple","pie","apple","red","red","red"}
	fmt.Println("Test case 0, most repeated word: \t", mostRepeatedWord(words))

	fmt.Println("----------------------------------------")

	// Test case 1
	words = []string{"deniz", "aydemir", "is", "a", "student"}
	fmt.Println("Test case 1, most repeated word: \t", mostRepeatedWord(words))

	fmt.Println("----------------------------------------")

	// Test case 2
	words = []string{"banana", "apple", "grape", "kiwi", "avocado", "papaya", "kiwi"}
	fmt.Println("Test case 2, most repeated word: \t", mostRepeatedWord(words))

	fmt.Println("----------------------------------------")

	// Test case 3
	words = []string{"aaa", "bbb", "ccc", "bbb"}
	fmt.Println("Test case 3, most repeated word: \t", mostRepeatedWord(words))

	fmt.Println("----------------------------------------")

	// Test case 4
	words = []string{"one", "one", "one", "two"}
	fmt.Println("Test case 4, most repeated word: \t", mostRepeatedWord(words))

}