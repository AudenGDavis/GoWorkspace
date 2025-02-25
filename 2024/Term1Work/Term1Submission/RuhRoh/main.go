package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("file")
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

	terminalReader := bufio.NewReader(os.Stdin)

	fmt.Println("Press Enter to get the truth: ")
	_, _ = terminalReader.ReadString('\n')
	fmt.Println(myRunes)
	fmt.Println("You now have the truth. ;^)")

}
