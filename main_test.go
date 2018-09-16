package main

import (
	"reflect"
	"testing"
)

func TestCase(t *testing.T) {
	arr := []int{90, 80, 50, 40, 70}
	sortedArr := []int{40, 50, 70, 80, 90}
	if !reflect.DeepEqual(bubbleSort(arr), sortedArr) {
		t.Errorf("Expected: %v, Output: %v\n", sortedArr, bubbleSort(arr))
	}
}
