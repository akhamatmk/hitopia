package main

import (
	"fmt"
)

// Function to calculate the weights of uniform substrings
func weightedUniformStrings(s string, queries []int) []string {
	weights := make(map[int]bool) // Map to store the weights of uniform substrings

	// Calculate weights
	i := 0
	for i < len(s) {
		weight := int(s[i] - 'a' + 1)
		count := 1
		weights[weight] = true

		// Increment count for consecutive characters with the same weight
		for i+1 < len(s) && s[i+1] == s[i] {
			count++
			weight += int(s[i+1] - 'a' + 1)
			weights[weight] = true
			i++
		}

		// Increment i to the next unique character
		i++
	}

	// Check if the given queries exist in the weights map
	result := make([]string, len(queries))
	for i, q := range queries {
		if weights[q] {
			result[i] = "Yes"
		} else {
			result[i] = "No"
		}
	}

	return result
}

func main() {
	s := "abccddde"
	queries := []int{1, 3, 12, 5, 9, 10}
	result := weightedUniformStrings(s, queries)

	fmt.Println(result)
	fmt.Println("Results:")
	for i, res := range result {
		fmt.Printf("Query %d: %s\n", queries[i], res)
	}
}
