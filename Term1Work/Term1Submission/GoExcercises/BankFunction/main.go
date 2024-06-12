package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Intial Value: ")
	initial, _ := reader.ReadString('\n')
	initial, _ = strings.CutSuffix(initial, "\n")
	initial2, _ := strconv.ParseFloat(initial, 64)

	fmt.Print("Enter Yearly Interest (%): ")
	interest, _ := reader.ReadString('\n')
	interest, _ = strings.CutSuffix(interest, "\n")
	interest2, _ := strconv.ParseFloat(interest, 64)

	fmt.Print("Enter Time: ")
	time, _ := reader.ReadString('\n')
	time, _ = strings.CutSuffix(time, "\n")
	time2, _ := strconv.ParseFloat(time, 64)

	fmt.Println(getInterest(initial2, interest2, time2))
}

func getInterest(initial, interest, time float64) float64 {
	return initial * math.Pow((interest)/100+1, time)
}
