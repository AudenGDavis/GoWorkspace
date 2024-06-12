package main

import "fmt"

func main() {
	fmt.Println(myFunction(2))
	myFunctionButNoReturn(2)
}

func myFunction(value int) int {
	return value * 2
}

func myFunctionButNoReturn(value int) {
	fmt.Println(value, ", but no print")
}
