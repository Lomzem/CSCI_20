/***************************************
Input Number / Output Number
Author: Lawjay Lee
Date Completed: 2023-08-28
Description: Gets user input, prints user input then prints double user input
***************************************/

package main

import "fmt"

func main() {
    var input int
    fmt.Println("Please enter an integer:")
    fmt.Scanln(&input)
    fmt.Println("You entered", input)
    fmt.Println("Twice that value is", input * 2)
}

