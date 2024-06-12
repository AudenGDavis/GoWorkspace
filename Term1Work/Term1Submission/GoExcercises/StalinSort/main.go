package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := []int{}

	for i := 0; i < 100; i++ {
		slice = append(slice, rand.Intn(100)+1)
	}

	fmt.Println(slice)

	result := []int{}

	largestIndex := 0

	for i := 0; i < len(slice); i++ {
		if slice[largestIndex] <= slice[i] {

			largestIndex = i
			result = append(result, slice[largestIndex])

		}
	}

	fmt.Println(result)
}
