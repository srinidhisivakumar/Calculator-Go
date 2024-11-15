package main

import (
	"fmt"
)

// Square function by Anjani
func Square(a int) int {
	return a * a
}

func main() {
	square := Square(5)
	fmt.Println("Square:", square)
}
