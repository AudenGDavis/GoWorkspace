package main

import (
	"fmt"
)

func main() {
	list := [50]int{}
	for i := 0; i < len(list); i++ {
		list[i] = i
	}

	fmt.Println(list)

	goal := 23
	lowerIndex := 0
	upperIndex := len(list) - 1
	for lowerIndex != upperIndex-1 {
		middle := (upperIndex - lowerIndex) / 2
		fmt.Printf("%d - %d - %d\n", lowerIndex, middle, upperIndex)
		if middle > goal {
			upperIndex = middle
		} else if middle < goal {
			lowerIndex = goal
		} else if
		

	}

	fmt.Printf("%d - %d\n", lowerIndex, upperIndex)

}
