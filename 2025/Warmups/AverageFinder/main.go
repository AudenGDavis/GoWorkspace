package main

import "fmt"

func main() {
	var sum float32 = 0
	list := []float32{1, 2, 3, 4, 5}
	for i := 0; i < len(list); i++ {
		sum += list[i]
	}
	fmt.Printf("Average: %f\n", (sum / float32(len(list))))
}
