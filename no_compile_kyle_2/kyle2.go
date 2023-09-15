/***************************************
No Compile Kyle 2
Author: Lawjay Lee
Date Completed: 2023-09-14
Description: Fix kyle's mistakes 2
***************************************/

/***********************************************************
Program Name: Simple Math Calculator
Program Author: Kyle NoCompile
Date Created: 9/12/16
Program Description:
	This program performs simple arithmetic calculations.
	The user enters numbers and the math operation to
	perform on those numbers. The program will then display
	the result of the operation.

Modified Date:
Modified Description:
***********************************************************/

package main

// OLD: import fmt
// Should have fmt in quotes
import "fmt"

// This is the main function (where the program begins)
// OLD: fun main[] {
// should declare function with func, functions should use parentheses not brackets
func main() {

	// Variables to hold local data
	var firstNum, secondNum int
    // OLD: var mathChoice() string
    // variable names should not have parentheses
    var mathChoice string
    // OLD: result int
    // should declare variable with var
    var result int
	
	// Call the showWelcome() function
    // OLD: showWelcome("welcome")
    // showWelcome does not accepts parameters
    showWelcome()
	
	// Call the getInteger() function (for the first integer)
	// and store the result in the "firstNum" variable
    // OLD: firstNum = GetInteger(true)
    // should be getInteger
	firstNum = getInteger(true)
	
	
	// Call the getMathChoice() function and store result in "mathChoice" variable
    // OLD: mathChoice = getMathChoice(42)
    // getMathChoice doesnt accept args
	mathChoice = getMathChoice()
	
	// Call validateMathChoice() function, passing it the user's math choice
	// and using the return value to decide what to do next
    // OLD: if (validateMathChoice()) {
    // validateMathChoice requires an arg
	if (validateMathChoice(mathChoice)) {
    // OLD: {
    // bracket should be on the same line as the if statement
		// Call the getInteger() function (for the second and subsequent integers)
		// and store the result in the "secondNum" variable			
		secondNum = getInteger(false)
		
		// Call the doMath() function and pass it all of the user input
		// and store the return value in the "result" variable.
        // OLD: result = doMath(firstNum secondNum mathChoice)
        // function arguments should be separated by commas
        result = doMath(firstNum, secondNum, mathChoice)
		
		// Call the showResult() function to show the result
		showResult(result)

	} else {
    // OLD: else {
    // else should be on the same line as the closing bracket
		// If the user chose an invalid math function...
        // OLD: fmt.Println(Not a valid math choice)
        // fmt.Println only accepts strings
        fmt.Println("Not a valid math choice")
	}

}


// This function shows a nice welcome message
func showWelcome() {
	fmt.Println("******************************************")
	fmt.Println("Welcome to the simple calculator program!")
    // OLD: fmt.Println("This program will do simple addition and"
    // Println should have a closing parenthesis
	fmt.Println("This program will do simple addition and")
	fmt.Println("subtraction. Math is fun, so enjoy!")
	fmt.Println("******************************************")
}

// This function gets integer input from the user
// OLD: func getUserIntegerInput() string {
// function should return integer
func getUserIntegerInput() int {
// OLD: {
// bracket should be on the same line as function definition
	var input int
    // OLD: fmt.Scanln(input)
    // should be &input
	fmt.Scanln(&input)
	return input
}

// This function asks the user for a math function
// and returns the user's input
// OLD: function getMathChoice() string {
// should be func instead of function
func getMathChoice() string {
	var choice string
	fmt.Println("\nPlease select a math function to perform (\"+\" = Addition, \"-\" = Subtraction):")
	fmt.Scanln(&choice)
    // OLD: ret choice
    // should be return not ret
    return choice
}

// this function asks the user for either the first integer
// or the second and returns the user's input
func getInteger(firstNumber bool) int {

	fmt.Print("\nPlease enter the ")
	
	// if the "firstNumber" variable is true, then this
	// is the first number being collected
	if (firstNumber) {
        // OLD: format.Print("first ")
        // Should be fmt
		fmt.Print("first ")
	
	} else { // Otherwise, it's the second number being collected
        // OLD: format.Print("second ")
        // Should be fmt
        fmt.Print("second ")
	}
	
	fmt.Println("integer:")
	
	// Call the getUserIntegerInput() function and return the return value from it
    // OLD: return = getUserIntegerInput()
    // return doesnt need =
	return getUserIntegerInput()
}


// This function validates the user's match function choice
// by returning a true for valid, and a false for invalid
func validateMathChoice(choice string) bool {
    // OLD: if (choice == "+" && choice == "-") {
    // choice can't be both + and -, should be or
	if (choice == "+" || choice == "-") {
		return true
	} else 
	{
		return false
	}
}

// This function adds two integers
// OLD: func doAddition(int1 int, int2 int) integer {
// integer should be declared with int
func doAddition(int1 int, int2 int) int {
	return int1 + int2
}

// This function subtracts the second integer
// parameter from the first integer parameter
// OLD: funky doSubtraction(int1 int, int2 int) int {
// functions declared with func not funky
func doSubtraction(int1 int, int2 int) int {
	return int1 - int2
}


// This function determines the result of the math
// operation requested by the user
func doMath(firstInt int, secondInt int, mathFunc string) int {
	
	// Initialize result to zero (0)
    // OLD: var result int == 0
    // variables should be declared with one =
	var result int = 0
	
	// If the math function is a "+", then call the
	// doAddition() function and store the return 
	// value in the "result" variable
	if mathFunc == "+" {

        // OLD: result = DoAddition(firstInt, secondInt)
        // function name is doAddition not DoAddition
		result = doAddition(firstInt, secondInt)

	// If the math function is a "-", then call the
	// doSubtraction() function and store the return 
	// value in the "result" variable
    // OLD: } if else mathFunc == "-" {
    // should be else if
	} else if mathFunc == "-" {
        // OLD: result = DoSubtraction(firstInt, secondInt)
        // function name is doSubtraction not DoSubtraction
        result = doSubtraction(firstInt, secondInt)
	}
	
	return result
}

// This function displays the result of a math operation
func showResult(result int) {
    // OLD: fmt.Println("\nThe result is" result)
    // should separate arguments with comma
	fmt.Println("\nThe result is", result)
}
