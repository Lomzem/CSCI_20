/***************************************
No Compile Kyle
Author: Lawjay Lee
Date Completed: 2023-08-28
Description: Fix Kyle's errors, program should do what kyle described
***************************************/

/***********************************************************
Program Name: Triangle Area Calculator
Program Author: Kyle NoCompile
Date Created: 8/1/2019
Program Description: 
	This program gets the dimensions of a triangle (base
	and height) from the end user and then calculates the
	area of the triangle. (Base * Height) / 2 
	
Modified Date:
Modified Description:
***********************************************************/

// pack mane
// pack mane should be "package main"
package main

// include "format"
// golang uses "import" and the library is "fmt"

import "fmt"

// function main()
// functions are declared using func
func main() {
	// var ba$e float32
    // $ isn't allowed in var names
    var base float32
	// float32 1height var
    // variable declaration is var <name> <type>
    // can't start var name with number
    var height float32

	var _answer float32

	// var answer message string = "The answer is";
    // golang doesn't end statements with ;
    // can't use spaces in variable names
    var answer_message string = "The answer is"

	// Greet the user
	// Println("Hello, Friend!")
    // should use the fmt library
    fmt.Println("Hello, Friend!")
	
	// Ask for the triangle base
	// Println("\nWhat is the length of the base of the triangle?"}
    // Can't close method with } use ) instead; also use fmt library to println
    fmt.Println("\nWhat is the length of the base of the triangle?")

	// fmt.GetInputFromUser(&ba$e)
    // correct var naming
    // method is Scanln
    fmt.Scanln(&base)

	// Ask for the triangle height
	// fmt.Scanln("\nWhat is the height of the triangle?")
	// fmt.Println(&1height)

    // fmt.Println should print the question, fmt.Scanln should collect user input and store it in a var
    fmt.Println("\nWhat is the height of the triangle?")
    fmt.Scanln(&height)

	// Reassuring message
	// fmt.Println["Thanks, I'm doing some math now..."]
    // Methods are called using () not []
    fmt.Println("Thanks, I'm doing some math now... ")

	// Do some math
	// ba$e * 1height = _answer
    // fixed var names; variable should be on the left and value should be on the right side of the =
    _answer = base * height
	
	// Display the area of the triangle
	// fmt.Println(answer message, "(_answer / 2)")
    // fixed answer message var name
    // (_answer / 2) should not be in quotes since it's an expession
    fmt.Println(answer_message, _answer / 2)
}
