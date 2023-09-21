/***************************************
Dinner For One
Author: Lawjay Lee
Date Completed: 2023-09-20
Description: todo
***************************************/

package main

import "fmt"

type Combo struct {
    Entree string
    Side string
    Drink string
    Size int32
    Price float32
}

func NewCombo(Entree string, Side string, Drink string, Size int32, Price float32) Combo {
    if Drink == "" {
        Drink = "Coke"
    }
    return Combo{Entree, Side, Drink, Size, Price}
}

func (c Combo) Display() {
    // todo
}

func main() {
    var userDrink string
    var hamburger Combo = NewCombo("Hamburger", "Fries")
}
