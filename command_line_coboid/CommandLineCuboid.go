/****************************************************************
Command Line Cuboid
Author: Lawjay Lee
Date Completed: 2023-10-11
Description: Calculates volume of cuboid using command line args
****************************************************************/

package main

import (
	"flag"
	"fmt"
)

// Implement Cuboid data type definition here
type Cuboid struct {
    Width int
    Length int
    Height int
}

// Implement pseudo-constructor function (NewCuboid) here
func NewCuboid(width int, length int, height int) Cuboid {
    var c Cuboid
    c.Width = width
    c.Length = length
    c.Height = height
    return c
}

// Implement Cuboid member function definitions here
func (c *Cuboid) Volume() int {
    return c.Width * c.Length * c.Height
}

func (c * Cuboid) Display() {
    fmt.Println("Width:", c.Width)
    fmt.Println("Length:", c.Length)
    fmt.Println("Height:", c.Height)
}

func main() {
	var inputWidth, inputLength, inputHeight int

	// Implement logic to map and assign command line argument values
	// to inputWidth, inputLength, and inputHeight variables here
    flag.IntVar(&inputWidth, "Width", 0, "Width of cuboid")
    flag.IntVar(&inputLength, "Length", 0, "Length of cuboid")
    flag.IntVar(&inputHeight, "Height", 0, "Height of cuboid")
    flag.Parse()

	var customCuboid Cuboid = NewCuboid(inputWidth, inputLength, inputHeight)
	customCuboid.Display()
	fmt.Println("\nVolume:", customCuboid.Volume())
}
