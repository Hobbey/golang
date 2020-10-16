package binarySearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var testCaseTable = []struct {
		arr []int // input array
		target int //input target
		expected int // expected output index
	} {
		// case1: value found
		{[]int{11, 14, 19, 20, 31, 45, 63, 70, 78}, 14, 1},
		// case2: value not found
		{[]int{11, 14, 19, 20, 31, 45, 63, 70, 78}, 71, -1},
		// case3: value not found, out of range, avoid for loop
		{[]int{11, 14, 19, 20, 31, 45, 63, 70, 78}, 888, -1},
		// case4: value not found, array not valid, avoid for loop
		{[]int{}, 70, -1},
	}

	for i, c := range testCaseTable{
		fmt.Printf("TEST CASE %v, Search %v in array %v\n", i+1, c.target, c.arr)
		actual := BinarySearch(c.arr, c.target)
		fmt.Println()
		if actual != c.expected {
			t.Errorf("BinarySearch(%v,%v) = %v; expected %v", c.arr, c.target, actual, c.expected)
		}
	}
}
