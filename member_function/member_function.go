/***************************************
Member Function
Author: Lawjay Lee
Date Completed: 2023-09-20
Description: Creates 3 combo objects and displays contents to terminal
***************************************/

package main

import "fmt"

type Combo struct {
    Entree string
    Side string
    Price float64
}

func NewCombo(Entree string, Side string, Price float64) Combo {
    return Combo{Entree, Side, Price}
}

func (c Combo) Display() {
    fmt.Println("Entree:", c.Entree)
    fmt.Println("Side:", c.Side)
    fmt.Print("Price: $", c.Price)
    fmt.Println()
}

func main() {
    var hamburger Combo = NewCombo("Hamburger", "Fries", 5.99)
    var burrito Combo = NewCombo("Burrito", "Rice", 4.99)
    var salad Combo = NewCombo("Salad", "Breadsticks", 4.99)

    fmt.Println("Combo 1:")
    hamburger.Display()

    fmt.Println("Combo 2:")
    burrito.Display()

    fmt.Println("Combo 3:")
    salad.Display()
}
