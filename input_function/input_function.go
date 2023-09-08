/***************************************
Input Function
Author: Lawjay Lee
Date Completed: 2023-09-07
Description: Ask user for 3 calendar dates and saves it to a CalendarDate object
***************************************/

package main

import "fmt"

type CalendarDate struct {
    Date int
    Month int
    Year int
}

func BuildCalendarDate() CalendarDate {
    var newCalendar CalendarDate

    fmt.Println("Enter Date:")
    fmt.Scanln(&newCalendar.Date)

    fmt.Println("Enter Month:")
    fmt.Scanln(&newCalendar.Month)

    fmt.Println("Enter Year:")
    fmt.Scanln(&newCalendar.Year)
    fmt.Println()

    return newCalendar
}

func main() {
    fmt.Println("Date 1:")
    _ = BuildCalendarDate()

    fmt.Println("Date 2:")
    _ = BuildCalendarDate()

    fmt.Println("Date 3:")
    _ = BuildCalendarDate()
}
