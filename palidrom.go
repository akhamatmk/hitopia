package main

import "fmt"

// Recursive function to find the highest palindrome by changing at most k digits
func highestPalindrome(str string, k int) string {
	if k < 0 {
		return "-1"
	}

	// Base case: if the string is empty or a palindrome, return the string
	if len(str) <= 1 || isPalindrome(str) {
		return str
	}

	// Recursive case:
	// Find the highest palindrome by replacing characters from left to right
	maxPalindrome := str

	for i := 0; i < len(str); i++ {
		for j := '9'; j >= rune(str[i]); j-- {
			newStr := str[:i] + string(j) + str[i+1:]
			palindrome := highestPalindrome(newStr, k-1)

			// Compare the current palindrome with the highest palindrome found so far
			if palindrome > maxPalindrome {
				maxPalindrome = palindrome
			}
		}

		// If k is exhausted, no need to check further replacements
		if k == 0 {
			break
		}
	}

	return maxPalindrome
}

// Function to check if a string is a palindrome
func isPalindrome(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	str := "3943"
	k := 1

	highest := highestPalindrome(str, k)
	if highest == "-1" {
		fmt.Println("-1")
	} else {
		fmt.Println("Highest Palindrome:", highest)
	}
}
