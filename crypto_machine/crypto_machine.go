/***************************************
Crypto Machine
Author: Lawjay Lee
Date Completed: 2023-11-29
Description: Allows you to encrypt message or decrypt message
***************************************/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rotor struct {
	Adder      int
	Multiplier int
}

type CryptoMachine struct {
	Rotors []Rotor
}

func NewRotor(multiplier int, adder int) Rotor {
	var r Rotor
	r.Multiplier = multiplier
	r.Adder = adder
	return r
}

func NewCryptoMachine() CryptoMachine {
	var c CryptoMachine
	c.Rotors = ReadRotorFile()
	return c
}

func ReadRotorFile() []Rotor {
	file, _ := os.Open("rotors.txt")
	defer file.Close()

	var rotorList []Rotor

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line string = scanner.Text()
		var lineSplit = strings.Split(line, " ")
		multiplier, _ := strconv.Atoi(lineSplit[0])
		adder, _ := strconv.Atoi(lineSplit[1])
		var rotor Rotor = NewRotor(multiplier, adder)
		rotorList = append(rotorList, rotor)
	}

	return rotorList
}

func (r Rotor) GetMultiplier() int {
	return r.Multiplier
}

func (r Rotor) GetAdder() int {
	return r.Adder
}

func (c CryptoMachine) GetRotors() []Rotor {
	return c.Rotors
}

func (r Rotor) Cipher(original int) int {
	return (original * r.GetMultiplier()) + r.GetAdder()
}

func (r Rotor) Decipher(original int) int {
	return (original - r.GetAdder()) / r.GetMultiplier()
}

func (c CryptoMachine) Encrypt(ascii int) int {
	for _, rotor := range c.GetRotors() {
		ascii = rotor.Cipher(ascii)
	}
	return ascii
}

func (c CryptoMachine) Decrypt(eInt int) string {
	var rotors []Rotor = c.GetRotors()
	// Loop in reverse
	for i := len(rotors) - 1; i >= 0; i-- {
		eInt = rotors[i].Decipher(eInt)
	}
	return string(rune(eInt))
}

func RunEncryption(c CryptoMachine, message string) {
	var encryptedMessage []int
	for _, char := range message {
		var ascii int = int(char)
		var encipheredInt int = c.Encrypt(ascii)
		encryptedMessage = append(encryptedMessage, encipheredInt)
	}
	for i := 0; i < len(encryptedMessage)-1; i++ {
		fmt.Print(strconv.Itoa(encryptedMessage[i]) + " ")
	}
	fmt.Println(encryptedMessage[len(encryptedMessage)-1])
}

func RunDecryption(c CryptoMachine, message string) {
	var decryptedMessage string

	var messageInts []string = strings.Split(message, " ")
	for _, strInt := range messageInts {
		eInt, _ := strconv.Atoi(strInt)
		var char string = c.Decrypt(eInt)
		decryptedMessage += char
	}
	fmt.Println(decryptedMessage)
}

func main() {
	var c CryptoMachine = NewCryptoMachine()

	var mode string
	var message string
	flag.StringVar(&mode, "mode", "", "encrypt|decrypt")
	flag.StringVar(&message, "message", "", "message to encrypt/decrypt")
	flag.Parse()

	if mode == "encrypt" {
		RunEncryption(c, message)
	} else if mode == "decrypt" {
		RunDecryption(c, message)
	}
}
