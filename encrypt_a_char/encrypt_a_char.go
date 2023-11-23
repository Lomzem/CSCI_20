/***************************************
Encrypt a Character
Author: Lawjay Lee
Date Completed: 2023-11-23
Description: Takes in command line args
to convert keyboard character using affine
cipher algorithm
***************************************/

package main

import (
	"flag"
	"fmt"
)

func main() {
	var character string
	flag.StringVar(&character, "character", "", "Single keyboard character that will be converted")

	var multiplier int
	flag.IntVar(&multiplier, "multiplier", 0, "Integer value for multipling affine cipher")

	var adder int
	flag.IntVar(&adder, "adder", 0, "Integer value for adding part of affine cipher")

	flag.Parse()

	if len(character) != 1 {
		panic("\"character\" arg should only be one character long")
	}

	var ascii int = int(character[0])
	var encrypted_val int = ascii*multiplier + adder
	fmt.Println("Encrypted value:", encrypted_val)
}
