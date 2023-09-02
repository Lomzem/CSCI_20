/***************************************
Cubic Objects
Author: Lawjay Lee
Date Completed: 2023-09-01
Description: Asks user to input dimensions for two rectangular cuboids. Then, prints volume.
***************************************/

package main

import "fmt"

type Cuboid struct {
    length int
    width int
    height int
}

func makeCuboid() (Cuboid) {
    var new_cube Cuboid

    fmt.Println("Enter the length:")
    fmt.Scanln(&new_cube.length)

    fmt.Println("Enter the width:")
    fmt.Scanln(&new_cube.width)

    fmt.Println("Enter the height:")
    fmt.Scanln(&new_cube.height)
    fmt.Println()

    return new_cube
}

func (cube Cuboid) calcVolume() (int) {
    return cube.length * cube.width * cube.height
}

func main() {
    fmt.Println("Cuboid 1:")
    var cube_1 Cuboid = makeCuboid()

    fmt.Println("Cuboid 2:")
    var cube_2 Cuboid = makeCuboid()

    fmt.Println("Cuboid 1 Volume:", cube_1.calcVolume())
    fmt.Println("Cuboid 2 Volume:", cube_2.calcVolume())
}

