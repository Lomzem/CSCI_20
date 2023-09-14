/**************************************
Book-A-Dwelling
Author: Lawjay Lee
Date Completed: 2023-09-14
Description: Allows a user to book a dwelling for a certain number of nights and calculates costs.
**************************************/

package main

import "fmt"

type Dwelling struct {
    Id int
    City string
    State string
    Bedrooms int
    Bathrooms int
    NightlyRate int
}

func NewDwelling(Id int, City string, State string, Bedrooms int, Bathrooms int, NightlyRate int) (Dwelling) {
    return Dwelling{Id, City, State, Bedrooms, Bathrooms, NightlyRate}
}

func (dwelling Dwelling) printInfo() {
    fmt.Println("Dwelling", dwelling.Id)
    fmt.Printf("Location: %s, %s\n", dwelling.City, dwelling.State)
    fmt.Printf("%d Bedrooms, %d Baths\n", dwelling.Bedrooms, dwelling.Bathrooms)
    fmt.Printf("Nightly Rate: $%d\n", dwelling.NightlyRate)
    fmt.Println()
}

func (dwelling Dwelling) calculateCosts(nights int) (int) {
    return dwelling.NightlyRate * nights
}

func main() {
    fmt.Println("Welcome to Book-A-Dwelling!")
    fmt.Println("\nHere are your dwelling options:")
    fmt.Println()

    var losAngeles Dwelling = NewDwelling(1, "Los Angeles", "CA", 4, 2, 300)
    var newYork Dwelling = NewDwelling(2, "New York", "NY", 3, 1, 245)
    var chicago Dwelling = NewDwelling(3, "Chicago", "IL", 4, 1, 200)
    var seattle Dwelling = NewDwelling(4, "Seattle", "WA", 5, 3, 375)

    dwellingIds := map[int]Dwelling{
        1: losAngeles,
        2: newYork,
        3: chicago,
        4: seattle,
    }

    losAngeles.printInfo()
    newYork.printInfo()
    chicago.printInfo()
    seattle.printInfo()

    fmt.Println("Please select a dwelling number from the list above")
    var idChoice int
    fmt.Scanln(&idChoice)
    var dwellingChoice Dwelling = dwellingIds[idChoice]

    fmt.Println("\nExcellent choice! How many nights will you stay?")
    var nightsChoice int
    fmt.Scanln(&nightsChoice)

    fmt.Println("\nBooking Summary:")
    dwellingChoice.printInfo()

    fmt.Println("Staying for", nightsChoice, "nights")
    fmt.Print("\nBase price of your stay: $", dwellingChoice.calculateCosts(nightsChoice))
}
