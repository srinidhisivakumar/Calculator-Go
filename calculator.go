package main

import (
	"errors"
	"fmt"
)

// Square function by Anjani
func Square(a int) int {
	return a * a
}

// Divide performs division and returns an error if dividing by zero.
func Divide_Case1(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Error: Can't be divided by zero")
	}
	return a / b, nil
}
func Divide_Case2(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Error: Can't be divided by zero")
	}
	return a / b, nil
}

func main() {
	square := Square(5)
	fmt.Println("Square:", square)

	result, error := Divide_Case1(5, 0)
	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Division:", result)
	}
	result, error = Divide_Case2(5, 5)
	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Division:", result)
	}
}
