package array

import "fmt"

func StaticArrayExample() {
	fmt.Println("Static array example")
	// A static array's space cannot be modified at runtime. The space is allocated in advance.
	arr := [3]int{1,2,3}

	for index, value := range arr {
		fmt.Printf("Index: %d; Value: %d\n", index, value)
	}

	arr[0] = 15

	fmt.Println(arr)
}