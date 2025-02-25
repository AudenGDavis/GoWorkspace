package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var length int = 20

	list := []int{}

	for i := 0; i < length; i++ {
		list = append(list, int(rand.Float32()*10+1))
	}

	fmt.Println("Originial: ")
	fmt.Println(list)

	for a := 0; a < len(list)-1; a++ {
		min := list[a]
		minIndex := a

		for b := a + 1; b < len(list); b++ {

			if list[b] < min {
				min = list[b]
				minIndex = b
			}
		}

		if minIndex != a {
			temp := list[a]
			list[a] = list[minIndex]
			list[minIndex] = temp
		}

	}
	fmt.Println("Final: ")
	fmt.Println(list)
}
