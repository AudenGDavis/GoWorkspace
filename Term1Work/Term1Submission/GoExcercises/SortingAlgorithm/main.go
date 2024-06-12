package main

import (
	"fmt"
	"math/rand"
)

func main() {

	arr := [10]int{}

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) + 1
	}

	for p := 0; p < len(arr); p++ {
		min := p
		for s := p; s < len(arr); s++ {
			if arr[min] > arr[s] {
				min = s
			}
		}

		if min != p {
			t := arr[p]

			arr[p] = arr[min]

			arr[min] = t
		}
	}

	fmt.Println(arr)
}
