/***************************************
Count Off
Author: Lawjay Lee
Date Completed: 2023-10-11
Description: Asks user to pick number from
1-10 then counts until number reaches 10
***************************************/

package main

import "fmt"

func main() {
	var userNum int
	fmt.Println("Enter a number between 1 and 10")
	fmt.Scanln(&userNum)
	fmt.Println()

	for i := userNum; i <= 10; i++ {
		fmt.Println(i)
	}

	var repeat string
	fmt.Println("Would you like to run the program again? (Y for yes, N for no)")
	fmt.Scanln(&repeat)

	if repeat == "Y" {
		main()
	} else {
		fmt.Println("\nOK, have a nice day!")
	}
}
