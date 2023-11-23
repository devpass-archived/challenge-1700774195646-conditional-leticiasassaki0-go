package main

import (
	"testing"
)

func TestCheckEvenOrOdd(t *testing.T) {
	result := CheckEvenOrOdd(10)
	if result != "even" {
		t.Errorf("Expected 'even' for input 10, but got %s", result)
	}

	result = CheckEvenOrOdd(15)
	if result != "odd" {
		t.Errorf("Expected 'odd' for input 15, but got %s", result)
	}
}
