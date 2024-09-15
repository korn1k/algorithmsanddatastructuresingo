package sort

import "fmt"

func mergeSort(arr []int) []int {
	// Base case: if the array has 1 or 0 elements, it's already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Find the middle index
	mid := len(arr) / 2

	// Recursively sort the left and right halves
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Merge the two sorted halves
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Compare elements from both arrays and add the smaller one to the result
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from left or right
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}


func MergeSortExample() {
	fmt.Println("Merge sort example")

	arr := []int{5,10,3,17,8}
	fmt.Println("Before sort: ", arr)
	arr = insertSort(arr)
	fmt.Println("After sort: ", arr)
}