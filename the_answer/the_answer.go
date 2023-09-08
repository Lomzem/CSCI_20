/***************************************
The Answer
Author: Lawjay Lee
Date Completed: 2023-09-07
Description: Asks user for an answer, then display whether they got it correct
***************************************/

package main

import "fmt"

func main() {
    fmt.Println("What is The Answer to the Ultimate Question of Life, the Universe, and Everything?")
    var userAnswer int
    fmt.Scanln(&userAnswer)

    if userAnswer == 42 {
        fmt.Println("You are correct")
    } else {
        fmt.Println("Nope")
    }
}
