package binarysearch

import "fmt"

func arrayBinarySearch(arr []int, value int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left) / 2

		if arr[mid] == value {
			return mid
		} else if arr[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func ArrayBinarySearchExample() {
	fmt.Println("Array binary search example")

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    value := 5
    index := arrayBinarySearch(arr, value)
    if index != -1 {
        fmt.Printf("Target %d found at index %d\n", value, index)
    } else {
        fmt.Printf("Target %d not found\n", value)
    }
}