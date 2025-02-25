package main

import (
	"fmt"
)

type Line struct {
	x1 float64
	y1 float64

	x2 float64
	y2 float64
}

func (line *Line) getSlope() float64 {
	return (line.y2 - line.y1) / (line.x2 - line.x1)
}

func (line *Line) getIntercept() float64 {
	return line.y1 - line.x1*line.getSlope()
}

func (line *Line) getRange() (float64, float64) {

	if line.y1 > line.y2 {
		return line.y2, line.y1
	}
	return line.y1, line.y2

}

func (line *Line) getDomain() (float64, float64) {

	if line.x1 > line.x2 {
		return line.x2, line.x1
	}
	return line.x1, line.x2

}

func newLine(X1 float64, Y1 float64, X2 float64, Y2 float64) *Line {
	returnthing := Line{X1, Y1, X2, Y2}
	return &returnthing
}

/*
IsIntersecting() Logic Chain
1 - Check if both lines are verticle and not touching. ==> return false

2 - Check if line1 is verticle, but line2 isn't. ==> Plug in line1.x1 into line2 and pass result

3 - Check if line2 is verticle, but line1 isn't. ==> Plug in line2.x1 into line1 and pass result

4 - Use Formula to find possible intercept
*/

func isIntersecting(line1, line2 Line) bool {
	var possibleX float64
	var possibleY float64
	if line1.x1 == line1.x2 && line2.x1 == line1.x2 && line1.x1 != line2.x1 { // Test Case 1
		return false
	} else if line1.x1 == line1.x2 { // Test Case 2
		possibleY = line1.x1 * line2.getSlope()
		possibleX = line1.x1
	} else if line2.x1 == line2.x2 { // Test Case 2
		possibleY = line2.x1 * line1.getSlope()
		possibleX = line2.x1
	} else {
		possibleX = (line2.getIntercept() - line1.getIntercept()) / (line1.getSlope() - line2.getSlope())
		possibleY = possibleX*line1.getSlope() + line1.getIntercept()
	}
	fmt.Println("(", possibleX, " , ", possibleY, ")")

	return false
}
