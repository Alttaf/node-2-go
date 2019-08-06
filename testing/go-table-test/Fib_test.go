package main

import (
	"fmt"
	"testing"
)

type tableTestData struct {
	input    int
	expected int
}

var tableTest = []tableTestData{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
}

// TestData this does some cool table testing stuff
func TestData(t *testing.T) {
	for _, tt := range tableTest {
		actual := Fib(tt.input)
		fmt.Printf("Fib(%d): resulted in %d \n", tt.input, actual)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.input, tt.expected, actual)
		}
	}

}
