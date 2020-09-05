func main() {
	items := []int{11, 14, 19, 20, 31, 45, 63, 70, 78}
	fmt.Println("index = ", binarySearch.BinarySearch(items, 12))
}


--------------------
0 4 8
12 <= arr[4], Target on Left: [0,4]
--------------------
0 2 4
12 <= arr[2], Target on Left: [0,2]
--------------------
0 1 2
12 <= arr[1], Target on Left: [0,1]
--------------------
0 0 1
arr[0] < 12, Target on Right: (0,1] ==> [1,1]
for loop Exit
**********************
Value in range but NOT found
index =  -1


