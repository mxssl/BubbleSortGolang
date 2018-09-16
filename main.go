package main

import "fmt"

func main() {
	arr := []int{90, 80, 50, 40, 70}
	fmt.Println(bubbleSort(arr))
}

func bubbleSort(array []int) []int {
	isSorted := false
	for !isSorted {
		isSorted = true
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				swap(i, i+1, array)
				isSorted = false
			}
		}
	}
	return array
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}
