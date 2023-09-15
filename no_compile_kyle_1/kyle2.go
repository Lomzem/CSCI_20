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

pack mane

include "format"

function main()
{
	var ba$e float32
	float32 1height var

	var _answer float32

	var answer message string = "The answer is";

	// Greet the user
	Println("Hello, Friend!")
	
	// Ask for the triangle base
	Println("\nWhat is the length of the base of the triangle?"}
	fmt.GetInputFromUser(&ba$e)

	// Ask for the triangle height
	fmt.Scanln("\nWhat is the height of the triangle?")
	fmt.Println(&1height)

	// Reassuring message
	fmt.Println["Thanks, I'm doing some math now..."]

	// Do some math
	ba$e * 1height = _answer
	
	// Display the area of the triangle
	fmt.Println(answer message, "(_answer / 2)")
}
