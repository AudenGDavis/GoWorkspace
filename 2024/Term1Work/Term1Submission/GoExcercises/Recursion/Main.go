package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter an int to find: ")
	stringInput, _ := reader.ReadString('\n')

	stringInput = strings.TrimSuffix(stringInput, "\n")

	input, _ := strconv.Atoi(stringInput)

	fmt.Println(GetFactorial(int64(input)))

}

func GetFactorial(i int64) int64 {
	if i <= 1 {
		return i
	}

	return GetFactorial(i-1) * i
}
