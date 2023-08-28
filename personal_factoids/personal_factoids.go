/***************************************
Personal Factoids
Author: Lawjay Lee
Date Completed: 2023-08-28
Description: Prints personal factoids about myself to stdout
***************************************/

package main

import "fmt"

func main() {
    fmt.Printf("Name: %s\n", "Lawjay Lee")
    fmt.Printf("Favorite Fictional Character: %s\n", "Kaneki Ken")
    fmt.Printf("Least Favorite Food: %s\n", "Gizzard")
    fmt.Printf("Star Trek or Star Wars: %s\n\n", "Star Trek")

    var favorite_quote string = "I am become death, the destroyer of worlds."

    fmt.Printf("Favorite Quote:\n%s", favorite_quote)
}

