/***************************************
Programmer's Dilemma
Author: Lawjay Lee
Date Completed: 2023-09-01
Description: Asks user to give name to colleague, then choose to assign blame to colleague, get outcome of situation
***************************************/

package main

import (
	"fmt"
	"math/rand"
    "time"
)

func init() {
    rand.Seed(int64(time.Now().Nanosecond()))
}

type Colleague struct {
    Name string
    Blame int
    Blamed bool
}

func makeColleague() Colleague {
    var colleague Colleague
    fmt.Println("What is your colleague's first name?")
    fmt.Scanln(&colleague.Name)

    // Makes Colleague.Blame random between true/false
    colleague.Blame = rand.Intn(2)
    colleague.Blamed = colleague.Blame == 1
    return colleague
}

func main() {
    var colleague = makeColleague()

    var blameDecision int
    fmt.Println("\nWill you blame", colleague.Name, "for the software bug? (0 = no, 1 = yes)")
    fmt.Scanln(&blameDecision)

    var blameColleague bool = blameDecision == 1

    fmt.Println()

    if blameColleague && colleague.Blamed {
        fmt.Println("You both blamed each other and so you both must work all weekend to fix the bug and take a significant pay cut.")
    } else if blameColleague && !colleague.Blamed {
        fmt.Println("You blamed", colleague.Name, "and they didn't blame you.", colleague.Name, "will be fired and you will get a promotion to lead a new team to fix the bug.")
    } else if !blameColleague && colleague.Blamed {
        fmt.Println(colleague.Name, "blamed you and you didn't blame them. You will be fired and", colleague.Name, "will get a promotion to lead a new team to fix the bug.")
    } else if !blameColleague && !colleague.Blamed {
        fmt.Println("Nobody blamed anyone. You and", colleague.Name, "will work this weekend to fix the bug.")
    }
}
