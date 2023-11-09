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
	"strings"
)

func main() {
	var shoppingList []string

	fmt.Println("Enter your shopping list. Enter * to indicate you are done:")
	var userInput string
	for userInput != "*" {
		fmt.Scanln(&userInput)
		shoppingList = append(shoppingList, userInput)
	}
	shoppingList = shoppingList[:len(shoppingList)-1]
	fmt.Println("Here is your list:")
	fmt.Println(strings.Join(shoppingList, ", "))
}
