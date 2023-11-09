/***************************************
Old School Encoder
Author: Lawjay Lee
Date Completed: 2023-11-08
Description: Accepts command line argument
for a number, converts the input from string
to int, then prints associated ASCII value.
***************************************/

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var InputNumber string
	flag.StringVar(&InputNumber, "InputNumber", "", "Number to convert to ASCII")
	flag.Parse()

	InputInt, err := strconv.Atoi(InputNumber)

	if err != nil {
		fmt.Println("InputNumber argument value must be an integer value")
		os.Exit(0)
	}

	if !(InputInt >= 0 && InputInt <= 127) {
		fmt.Println("InputNumber argument value must be an integer between 0 and 127")
		os.Exit(0)
	}

	fmt.Println(string(InputInt))
}
