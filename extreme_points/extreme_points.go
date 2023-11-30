/***************************************
Extreme Points
Author: Lawjay Lee
Date Completed: 2023-11-29
Description: Reads cartesian points from file,
prints min and max x and y values
***************************************/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func NewPoint(X int, Y int) Point {
	var point Point
	point.X = X
	point.Y = Y
	return point
}

func (p Point) GetX() int {
	return p.X
}

func (p Point) GetY() int {
	return p.Y
}

func (p *Point) GetCoord() string {
	var x string = strconv.Itoa(p.GetX())
	var y string = strconv.Itoa(p.GetY())
	return "(" + x + ", " + y + ")"
}

func GetPoints() []Point {
	file, _ := os.Open("points.txt")
	defer file.Close()

	var pointsList []Point

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line string = scanner.Text()
		var coords []string = strings.Split(line, ",")
		xValue, _ := strconv.Atoi(coords[0])
		yValue, _ := strconv.Atoi(coords[1])
		var point Point = NewPoint(xValue, yValue)
		pointsList = append(pointsList, point)
	}
	return pointsList
}

func main() {
	var pointsList []Point = GetPoints()
	fmt.Println(len(pointsList), "coordinate points found:")
	fmt.Println()

	var minXCoord Point = pointsList[0]
	var minYCoord Point = pointsList[0]
	var maxXCoord Point = pointsList[0]
	var maxYCoord Point = pointsList[0]

	for _, point := range pointsList {
		if point.GetX() < minXCoord.GetX() {
			minXCoord = point
		}
		if point.GetY() < minYCoord.GetY() {
			minYCoord = point
		}
		if point.GetX() > maxXCoord.GetX() {
			maxXCoord = point
		}
		if point.GetY() > maxYCoord.GetY() {
			maxYCoord = point
		}

		fmt.Println(point.GetCoord())
	}

	fmt.Println("Minimum X Coordinate Point:", minXCoord.GetCoord())
	fmt.Println("Maximum X Coordinate Point:", maxXCoord.GetCoord())
	fmt.Println("Minimum Y Coordinate Point:", minYCoord.GetCoord())
	fmt.Println("Maximum Y Coordinate Point:", maxYCoord.GetCoord())
}
