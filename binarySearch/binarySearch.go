package binarySearch

import "fmt"

func BinarySearch(arr []int, val int) (index int) {
	// todo: 增加数据是否顺序的检查

	// ....(负无穷).... 0 1 2 3 4 5 6 7 ....(正无穷)....
	// 先排除掉 从"有限数组arr"的两端 到"-∞"和"+∞"之间的广阔空间
	low := 0
	high := len(arr) - 1
	if len(arr) == 0 || arr[low] > val || arr[high] < val {
		fmt.Printf("Array not valid or given value out of range, return index = -1\n")
		return -1
	}

	// value 可能存在 [low,high] 的闭区间
	// low ----  mid ---- high
	// mid < value 时, value 可能存在 (mid,high] 的开闭区间, 也就是 [mid+1,high] 的闭区间
	// 否则 , 就是 value <= mid, 也就是 value 可能存在 [low,mid] 的闭区间
	for low < high {
		fmt.Println("--------------------")
		mid := (low + high) / 2
		fmt.Println(low, mid, high)
		if arr[mid] < val {
			fmt.Printf("arr[%v] < %v, Target on Right: (%v,%v] ==> [%v,%v]\n", mid, val, mid, high, mid+1, high)
			low = mid + 1
		} else {
			fmt.Printf("%v <= arr[%v], Target on Left: [%v,%v]\n", val, mid, low, mid)
			high = mid
		}
	}
	fmt.Println("for loop Exit")

	// low = high 时, 退出循环, 确定到数组内一个唯一位置
	// 此时 要么 arr[low] == val , value found , return index
	// 要么 arr[low] != val , value not found, return -1
	fmt.Println("**********************")
	if arr[low] == val {
		fmt.Printf("Value found, return index = %v\n", low)
		return low
	} else {
		fmt.Printf("Value in range but NOT found, return index = -1\n")
		return -1
	}
}
