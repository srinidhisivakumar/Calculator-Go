package main

import "testing"

func TestDivide(t *testing.T) {
	result := Divide(6, 2)
	expected := 3
	if result != expected {
		t.Fatalf("Divide(6, 2) will result %d, but we got %d", result, expected)
	}
}

func TestSquare (t *testing.T) {
	result := sqaure(2)
	expected := 4
	if result != expected {
		t.Fatalf("square(2) will result %d, but we got %d", result, expected)
	}
}