package main

import (
	"fmt"
	"sort"
	"strings"
)

// Function to sort words by the number of char a's within each word in descending order.
// If two words have the same number of a's, they will be sorted by length in descending order.
// If two words have the same number of a's and the same length, they will not be sorted.
// Input: words []string
// Output: None (words will be sorted in-place)
func sortWordsByA(words []string) {
	// A funtion to declare a custom comparator for the sort function
	comparator := func(i, j int) bool {
		// Get the number of char a's in each word
		countA1 := strings.Count(words[i], "a")
		countA2 := strings.Count(words[j], "a")

		// If the number of char a's in each word is the same, sort by length
		if countA1 == countA2 {
			// If the length of the two words are the same, do not sort
			if len(words[i]) == len(words[j]) {
				return false
			}
			return len(words[i]) > len(words[j])
		}
		return countA1 > countA2
	}

	// Sort the words by the custom comparator
	// As the sort function is in-place, the words slice will be sorted
	sort.Slice(words, comparator)
}

func main() {
	// Test case 0
	words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	fmt.Println("Test case 0, before sorting: \t", words)
	sortWordsByA(words)
	fmt.Println("Test case 0, after sorting: \t", words)

	fmt.Println("----------------------------------------")

	// Test case 1
	words = []string{"deniz", "aydemir", "is", "a", "student", "at", "the", "bilkent", "university", "in", "ankara", "turkey"}
	fmt.Println("Test case 1, before sorting: \t", words)
	sortWordsByA(words)
	fmt.Println("Test case 1, after sorting: \t", words)

	fmt.Println("----------------------------------------")

	// Test case 2
	words = []string{"banana", "apple", "grape", "kiwi", "avocado", "papaya"}
	fmt.Println("Test case 2, before sorting: \t", words)
	sortWordsByA(words)
	fmt.Println("Test case 2, after sorting: \t", words)

	fmt.Println("----------------------------------------")

	// Test case 3
	words = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Println("Test case 3, before sorting: \t", words)
	sortWordsByA(words)
	fmt.Println("Test case 3, after sorting: \t", words)

	fmt.Println("----------------------------------------")

	// Test case 4
	words = []string{"aaa", "abc", "def", "cba", "bca", "aba", "bac", "ada", "ccd", "bcb"}
	fmt.Println("Test case 4, before sorting: \t", words)
	sortWordsByA(words)
	fmt.Println("Test case 4, after sorting: \t", words)
}