package main

import "fmt"

func merge(left, right []int) []int {
	result := make([]int, 0)

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Add remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	start, end := 0, len(arr)

	mid := start + (end-start)/2

	left := mergeSort(arr[start:mid])
	right := mergeSort(arr[mid:end])

	return merge(left, right)
}

func main() {
	val := []int{3, 1, 2, 4, 1, 5, 2, 6, 4}

	fmt.Println(mergeSort(val))
}
