package binarysearch

import "fmt"

// Function to find the left boundary (first occurrence)
func findLeftBoundary(arr []int, target int) int {
    left, right := 0, len(arr)-1
    result := -1
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] < target {
            left = mid + 1
        } else if arr[mid] > target {
            right = mid - 1
        } else {
            result = mid  // potential first occurrence
            right = mid - 1  // continue searching in the left half
        }
    }
    return result
}

// Function to find the right boundary (last occurrence)
func findRightBoundary(arr []int, target int) int {
    left, right := 0, len(arr)-1
    result := -1
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] < target {
            left = mid + 1
        } else if arr[mid] > target {
            right = mid - 1
        } else {
            result = mid  // potential last occurrence
            left = mid + 1  // continue searching in the right half
        }
    }
    return result
}

// Function to find the range of the target
func searchRange(arr []int, target int) []int {
    left := findLeftBoundary(arr, target)
    if left == -1 {
        return []int{-1, -1}  // target not found
    }
    right := findRightBoundary(arr, target)
    return []int{left, right}
}

func RangeBinarySearchExample() {
	fmt.Println("Range binary search example")

    arr := []int{5, 7, 7, 8, 8, 10}
    target := 8
    result := searchRange(arr, target)
    fmt.Printf("Range of target %d: %v\n", target, result)
}