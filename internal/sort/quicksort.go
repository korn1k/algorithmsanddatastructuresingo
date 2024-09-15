package sort

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		// Base case: arrays with 1 or 0 elements are already sorted
		return arr
	}

	// Partition the array and get the pivot index
	pivotIndex := partition(arr)

	// Recursively sort the elements before and after the pivot
	quickSort(arr[:pivotIndex])    // Left of the pivot
	quickSort(arr[pivotIndex+1:])  // Right of the pivot

	return arr
}

// partition rearranges the array and returns the pivot index
func partition(arr []int) int {
	pivot := arr[len(arr)-1]  // Choose the last element as the pivot
	i := 0                    // Index of the smaller element

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			// Swap arr[i] and arr[j] to move smaller elements to the left of the pivot
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// Place the pivot element at the correct position
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]

	return i  // Return the index of the pivot
}

func QuickSortExample() {
	fmt.Println("Quick sort example")

	arr := []int{5,10,3,17,8}
	fmt.Println("Before sort: ", arr)
	arr = quickSort(arr)
	fmt.Println("After sort: ", arr)
}