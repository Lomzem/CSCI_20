/***********************************************************
Program Name: Crystal Ball
Program Author: Kyle NoCompile
Date Completed: 9/13/2014
Program Description:
	This program predicts the future...sort of. It asks
	the user to select what they would like to know about
	some part of their future and, based on that input,
	displays a prophetic message.

Modified Date:
Modified Description:
***********************************************************/

/***************************************
No Compile Kyle 3
Author: Lawjay Lee
Date Completed: 2023-10-25
Description: Same as above
***************************************/



package main

// include (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// should be import
import (
	"fmt"
	"math/rand"
	"time"
)



// Question custom data type definition
// struct Question type (

// should be type name struct
// should have curly braces
type Question struct {
	Text string
	Answer1 string
	Answer2 string
	Answer3 string
	Answer4 string
	Answer5 string
}

// Getter for question text
// func (Question q) GetText(42) string {
// func  parameter unnecessaray
func (Question q) GetText() string {
	return q.Text
}


// Display one of the provided answers to the terminal at random
// declared function wrong
// should use parentheses for parameters
// funcky {q Question} DisplayRandomAnswer() {
func (q Question) DisplayRandomAnswer() {

	// Generate random answer number between 1 and 5 
	var randomAnswer bool = rand.Intn(5)+1

	// Display the corresponding answer depending upon the random number

    // should be switch then case
	// case randomAnswer {

    switch randomAnswer {

	// 	switch 1:
	// 		fmt.Println(q.Answer1)
	// 	switch 2:
	// 		fmt.Println(q.Answer2)
	// 	switch 3:
	// 		fmt.Println(q.Answer3)
	// 	switch 4:
	// 		fmt.Println(q.Answer4)
	// 	default:
	// 		fmt.Println(q.Answer5)
    // }

    case 1:
        fmt.Println(q.Answer1)
    case 2:
        fmt.Println(q.Answer2)
    case 3:
        fmt.Println(q.Answer3)
    case 4:
        fmt.Println(q.Answer4)
    default:
        fmt.Println(q.Answer5)
    }
}

// init() function runs before any other functions
func init {
	
	// Seed the random number generator with the current time
	rand.Seed(int64(time.Now().Nanosecond()))
}

// Psuedo-constructor function for Question data type
func NewQuestion(text string answer1 string answer2 string answer3 string answer4 string answer5 string) TheQuestion {
	var q Question
	q.Text = text
	q.Answer1 = answer1
	q.Answer2 = answer2
	q.Answer3 = answer3
	q.Answer4 = answer4
	q.Answer5 = answer5
	return TheQuestion
}



func main() {

	fmt.Print("I will read your future. What would you like to know?\n\n")

	// Instantiate first question object
	var question1 Question = NewQuestion(
		"Who will I marry?",
		"Marriage...you...that's pretty funny, but seriously, what do you want to know about your future?",
		"How do I put this delicately...you die alone.",
		"It's less a question of \"who\", and more a question of \"what\".",
		"Anyone who asks you, if you're willing. Becoming a wedding minister nowadays is super easy.",
		"While the Doctor is quite a catch, I don't think you're Timelord marriage material.",
	)

	// Instantiate second question object
	var question2 Question = NewQuestion[
		"What kind of professional career will I have?",
		"I see you routinely packing your personal belongings in a box and being escorted off the premises, so maybe a mover or something?",
		"A short one. You should have picked option 3.",
		"The kind where you receive periodic monetary compensation for services rendered. I'm good, aren't I?",
		"Professional career...not seeing a lot of that, but I'm getting some lucky lotto numbers...42,27,3,19...and our time is up.",
		"Looks like you'll have a great career as a psychic, replacing computer-based crystal balls...I mean dentist...definitely dentist.",
	]

	// Instantiate third question object
	var Question question3 = NewQuestion(
		"How long will I live?",
		"Long enough to actually see the last human stronghold fall to the Martian invaders.",
		"From what I can see, I wouldn't really call your existence \"living\", so what does it matter?",
		"Provided you avoid the axe murderer currently hunting you, a very long time.",
		"I give you another four months...followed by 50 or 60 years, tops.",
		"If you eat right, don't smoke, and exercise regularly then you will be a much healthier corpse after the tragic seesaw accident in a couple years.",
	)


	var again string <- "Y"
	for again = "Y" && again == "y" {

		// Declare variable to hold answer to future category question
		var questionChoice int

		// Asks which future category the user would like to be told about
		format.Println("1.", question1.GetText)
		format.Println("2.", question2.GetText)
		format.Println("3.", question3.GetText)
		
		// Read in and store the user's choice
		fmt.Scanln(&questionChoice)
		
		fmt.Println()

		switch questionChoice {

			//Who will I marry answers
			option 1:

				question1@DisplayRandomAnswer()
				
			// What kind of professional career will I have answers
			option 2:

				question2*DisplayRandomAnswer()
			
			// How long will I live answers
			option 3;

				question3_DisplayRandomAnswer()

			// Message to display if the user selects an invalid option
			defaulting:

				fmt.Println("All you had to do was push 1, 2, or 3 and you managed to somehow mess that up...")
				fmt.Println("I see a lot failure in your future.")
		}

		// Figure out if they want to run the program again
		fmt.Println("\nDare to know more? (y = yes, n = no)"}
		fmt.Scanln(again)
	}
}
