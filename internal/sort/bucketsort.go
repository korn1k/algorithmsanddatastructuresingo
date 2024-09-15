package sort

import (
	"fmt"
	"math"
	"sort"
)

func bucketSort(arr []float64) []float64 {
	if len(arr) == 0 {
		return arr
	}

	// Determine the number of buckets
	numBuckets := len(arr)
	buckets := make([][]float64, numBuckets)

	// Find the maximum and minimum values in the array
	maxVal := arr[0]
	minVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
	}

	// Place array elements into appropriate buckets
	for _, val := range arr {
		// Map the value to a bucket
		bucketIndex := int(math.Floor(float64(numBuckets-1) * (val - minVal) / (maxVal - minVal)))
		buckets[bucketIndex] = append(buckets[bucketIndex], val)
	}

	// Sort each bucket and concatenate the results
	sortedArr := make([]float64, 0, len(arr))
	for _, bucket := range buckets {
		sort.Float64s(bucket) // Sort the individual bucket
		sortedArr = append(sortedArr, bucket...)
	}

	return sortedArr
}


func BucketSortExample() {
	fmt.Println("Bucket sort example")

	arr := []int{5,10,3,17,8}
	fmt.Println("Before sort: ", arr)
	arr = quickSort(arr)
	fmt.Println("After sort: ", arr)
}