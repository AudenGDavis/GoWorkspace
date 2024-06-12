package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	wallChars []rune = getRuneFile("./wallChars.txt")
)

type game struct {
	player  unit
	gameMap [][]rune
}

type unit struct {
	x int
	y int
}

func main() {
	myGame := game{}
	myGame.getFile("./map.txt")
	myGame.player = unit{1, 2}

	myGame.PrintGameInfo()
}

// print the 2d array with no border
func (myGame *game) PrintArrayNoBorder() {

	for y := 0; y < len(myGame.gameMap); y++ {
		for x := 0; x < len(myGame.gameMap[0]); x++ {
			fmt.Print(string(myGame.gameMap[y][x]))
		}
		fmt.Println()

	}
}

func (myGame *game) PrintGameInfo() {

	fmt.Print("+")
	for i := 0; i < len(myGame.gameMap[0]); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")

	for y := 0; y < len(myGame.gameMap); y++ {
		fmt.Print("|")
		for x := 0; x < len(myGame.gameMap[0]); x++ {
			fmt.Print(string(myGame.gameMap[y][x]))
		}
		fmt.Print("|\n")

	}

	fmt.Printf("|Player X: %-"+strconv.Itoa(len(myGame.gameMap[0])-10)+"d|\n", myGame.player.x)
	fmt.Printf("|Player Y: %-"+strconv.Itoa(len(myGame.gameMap[0])-10)+"d|\n", myGame.player.y)

	fmt.Print("+")
	for i := 0; i < len(myGame.gameMap[0]); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

// finds a file with the given file name and returns the 2d rune array
func (myGame *game) getFile(fileName string) {
	file, _ := os.Open(fileName)
	reader := bufio.NewReader(file)
	runeVal, _, _ := reader.ReadRune()
	returnArr := [][]rune{}
	returnArr = append(returnArr, []rune{})
	y := 0
	for i := 0; i < reader.Size() && runeVal != '!'; i++ {
		if runeVal != 0 && runeVal != rune(10) {
			returnArr[y] = append(returnArr[y], runeVal)
			runeVal, _, _ = reader.ReadRune()
		} else if runeVal == rune(10) {
			returnArr = append(returnArr, []rune{})
			y++
			runeVal, _, _ = reader.ReadRune()
		}

	}
	defer file.Close()

	myGame.gameMap = returnArr
}

func getRuneFile(fileName string) []rune {
	file, _ := os.Open(fileName)
	reader := bufio.NewReader(file)
	runeVal, _, _ := reader.ReadRune()
	returnArr := []rune{}
	for i := 0; i < reader.Size(); i++ {
		if runeVal != 0 && runeVal != rune(10) {
			returnArr = append(returnArr, runeVal)
			runeVal, _, _ = reader.ReadRune()
		}

	}
	defer file.Close()
	return returnArr
}

// print the 2d array with no border
func (myGame *game) PrintArray() {

	fmt.Print("+")
	for i := 0; i < len(myGame.gameMap[0]); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")

	for y := 0; y < len(myGame.gameMap); y++ {
		fmt.Print("|")
		for x := 0; x < len(myGame.gameMap[0]); x++ {
			fmt.Print(string(myGame.gameMap[y][x]))
		}
		fmt.Print("|\n")

	}

	fmt.Print("+")
	for i := 0; i < len(myGame.gameMap[0]); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func (myGame *game) makeMove(x, y int) bool {
	if myGame.player.y+y < 0 || myGame.player.y+y >= len(myGame.gameMap) {
		return false
	}
	if myGame.player.x+x < 0 || myGame.player.x+x >= len(myGame.gameMap[0]) {
		return false
	}
	for i := 0; i < len(wallChars); i++ {
		if myGame.gameMap[myGame.player.y+y][myGame.player.x+x] == wallChars[i] {
			return false
		}
	}

	myGame.player.x = x
	myGame.player.y = y

	return true
}
