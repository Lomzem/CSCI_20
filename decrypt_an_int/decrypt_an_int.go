/***************************************
Decrypt An Integer
Author: Lawjay Lee
Date Completed: 2023-11-23
Description: Decrypts resultant integer derived
from a character encrypted using affine cipher
***************************************/

package main

import (
	"flag"
	"fmt"
)

func main() {
	var encryptedInt int
	flag.IntVar(&encryptedInt, "encryptedInt", 0, "Integer value resulting from a previous encryption")

	var multiplier int
	flag.IntVar(&multiplier, "multiplier", 0, "Integer value from multipling part of previous encryption")

	var adder int
	flag.IntVar(&adder, "adder", 0, "Integer value from adding part of previous encryption")

	flag.Parse()

	var originalAscii int = (encryptedInt - adder) / multiplier
	var originalChar string = string(rune(originalAscii))
	fmt.Println("Decrypted value:", originalChar)
}
