package sort

import "fmt"

func insertSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		current := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = current
	}

	return arr
}

func InsertSortExample() {
	fmt.Println("Insert sort example")

	arr := []int{5,10,3,17,8}
	fmt.Println("Before sort: ", arr)
	arr = insertSort(arr)
	fmt.Println("After sort: ", arr)
}