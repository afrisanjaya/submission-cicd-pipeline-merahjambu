package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 1)
	expected := 2
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(2, 2)
	expected := 0
	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 3)
	expected := 9
	if result != expected {
		t.Errorf("Multiply(4, 3) = %d; want %d", result, expected)
	}
}
