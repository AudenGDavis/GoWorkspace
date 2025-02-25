package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(getFile("./test.txt"))
}

func getFile(fileName string) []rune {
	file, _ := os.Open(fileName)
	reader := bufio.NewReader(file)
	runeVal, _, _ := reader.ReadRune()
	myRunes := []rune{}
	for i := 0; i < reader.Size(); i++ {
		if runeVal != 0 {
			myRunes = append(myRunes, runeVal)
		}
		runeVal, _, _ = reader.ReadRune()
	}
	defer file.Close()
	return myRunes
}
