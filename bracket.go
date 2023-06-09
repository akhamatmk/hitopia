package main

import "fmt"

// Stack represents a stack data structure
type Stack struct {
	items []rune
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() rune {
	if len(s.items) == 0 {
		return 0
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// IsBalanced checks if the given string has balanced brackets
func IsBalanced(input string) bool {
	stack := Stack{}

	for _, char := range input {
		switch char {
		case '(', '[', '{':
			stack.Push(char)
		case ')':
			if stack.IsEmpty() || stack.Pop() != '(' {
				return false
			}
		case ']':
			if stack.IsEmpty() || stack.Pop() != '[' {
				return false
			}
		case '}':
			if stack.IsEmpty() || stack.Pop() != '{' {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

func main() {
	// Test cases
	expressions := []string{
		"{()}",
		"{(})",
		"[(){}]",
		"([]{}[[]])",
		"[(])",
		"{(([])[])[]}",
	}

	for _, expr := range expressions {
		if IsBalanced(expr) {
			fmt.Printf("%s is balanced\n", expr)
		} else {
			fmt.Printf("%s is not balanced\n", expr)
		}
	}
}
