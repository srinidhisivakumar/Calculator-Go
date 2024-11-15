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
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Error: Can't be divided by zero")
	}
	return a / b, nil
}

func main() {
	square := Square(5)
	fmt.Println("Square:", square)

	result, err := Divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result)
	}

}
