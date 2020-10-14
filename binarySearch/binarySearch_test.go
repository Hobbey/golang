package binarySearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var testTable = []struct {
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

	for i, testcase := range testTable{
		fmt.Printf("TEST CASE %v, Search %v in array %v\n", i+1, testcase.target, testcase.arr)
		actual := BinarySearch(testcase.arr, testcase.target)
		fmt.Println()
		if actual != testcase.expected {
			t.Errorf("BinarySearch(%v,%v) = %v; expected %v", testcase.arr, testcase.target, actual, testcase.expected)
		}
	}
}
