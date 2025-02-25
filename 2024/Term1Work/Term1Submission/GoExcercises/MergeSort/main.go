package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := [10]int{}

	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100 + 1)
	}

	fmt.Println(array, " => ", MergeSort(array[:]))
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	arr1 := []int{}
	arr2 := []int{}

	for i := 0; i < len(arr)/2; i++ {
		arr1 = append(arr1, arr[i])
	}

	for i := len(arr) / 2; i < len(arr); i++ {
		arr2 = append(arr2, arr[i])
	}

	arr1 = MergeSort(arr1)
	arr2 = MergeSort(arr2)

	returnArr := []int{}
	one := 0
	two := 0

	for one+two < len(arr) {

		if one >= len(arr1) {
			returnArr = append(returnArr, arr2[two])
			two += 1

			continue
		}

		if two >= len(arr2) {
			returnArr = append(returnArr, arr1[one])
			one += 1

			continue
		}

		if arr1[one] < arr2[two] {
			returnArr = append(returnArr, arr1[one])
			one += 1

			continue
		}

		if arr2[two] <= arr1[one] {
			returnArr = append(returnArr, arr2[two])
			two += 1
		}

	}

	return returnArr
}
