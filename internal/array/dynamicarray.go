package array

import "fmt"

func DynamicArrayExample() {
	fmt.Println("Dynamic array example")
	// A dynamic array's space can be modified at runtime.
	arr := []int{1,2,3,4}

	for index, value := range arr {
		fmt.Printf("Index: %d; Value: %d\n", index, value)
	}

	arr[0] = 4
	arr = append(arr, 8)

	fmt.Println(arr)
}