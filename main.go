package main

import "fmt"
import "golang/binarySearch"

func main() {
	items := []int{11, 14, 19, 20, 31, 45, 63, 70, 78}
	fmt.Println("index = ", binarySearch.BinarySearch(items, 12))
}