/***************************************
Shopping List
Author: Lawjay Lee
Date Completed: 2023-11-08
Description: Allows user to input list of items,
then prints list to terminal
***************************************/

package main

import (
	"fmt"
	// for standard library solution:
	// "strings"
)

func main() {
	var shoppingList []string

	fmt.Println("Enter your shopping list. Enter * to indicate you are done:")

	var userInput string
	for userInput != "*" {
		fmt.Scanln(&userInput)
		shoppingList = append(shoppingList, userInput)
	}

	fmt.Println("Here is your list:")

	for i := 0; i < len(shoppingList)-2; i++ {
		// Minus two is because the asterisk is appended to the list,
		// and we don't want to add a comma for the last item
		fmt.Print(shoppingList[i], ", ")
	}

	// len() - 1 would be asterisk, len() - 2 is last item
	fmt.Print(shoppingList[len(shoppingList)-2])

	// Messes with Scanln if there's no newline
	fmt.Println()

	// Solution with standard library:
	// fmt.Println(strings.Join(shoppingList, ", "))
}
