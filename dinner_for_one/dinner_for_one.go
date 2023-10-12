/***************************************
Dinner For One
Author: Lawjay Lee
Date Completed: 2023-09-20
Description: Asks user to select combo, drink, size, then calculates price
***************************************/

package main

import "fmt"

type Combo struct {
	Entree string
	Side   string
	Drink  string
	Size   int
	Price  float64
}

func NewCombo(Entree string, Side string, Drink string, Size int, Price float64) Combo {
	var combo Combo
	combo.Entree = Entree
	combo.Side = Side
	combo.Drink = Drink
	combo.Size = Size
	combo.Price = Price

	return combo
}

func (c *Combo) Display() {
	fmt.Println("Entree:", c.Entree)
	fmt.Println("Side:", c.Side)
	fmt.Printf("Price: $%g\n", c.Price)
}

func (c *Combo) SetSize(size int) {
	c.Size = size
	sizePrice := map[int]float64{
		1: 0,
		2: 2,
		3: 3,
	}
	c.Price += sizePrice[size]
}

func (c *Combo) SetDrink(drink string) {
	if drink != "" {
		c.Drink = drink
	}
}

func (c *Combo) DisplayFull() {
	sizeName := map[int]string{
		1: "Small",
		2: "Medium",
		3: "Large",
	}

	fmt.Println("Entree:", c.Entree)
	fmt.Println("Side:", c.Side)
	fmt.Println("Size:", sizeName[c.Size])
	fmt.Println("Drink:", c.Drink)
	fmt.Printf("Price: $%g", c.Price)
}

func PromptMenuAndSelectCombo() Combo {
	var hamburger Combo = NewCombo("Hamburger", "Fries", "Coke", 1, 5.99)
	var burrito Combo = NewCombo("Burrito", "Rice", "Coke", 1, 4.99)
	var salad Combo = NewCombo("Salad", "Breadsticks", "Coke", 1, 4.49)
	menuItems := []Combo{hamburger, burrito, salad}

	for i := 1; i <= len(menuItems); i++ {
		fmt.Printf("\nCombo %d:\n", i)
		menuItems[i-1].Display()
	}

	var orderNum int
	fmt.Println("\nPlease select your order number")
	fmt.Scanln(&orderNum)
	return menuItems[orderNum-1]
}

func UpdateSizeAndDrink(c Combo) Combo {
	var sizeChoice int
	fmt.Println("\nSize upgrade prices: Small $0, Medium $2, Large $3")
	fmt.Println("What size would you like? (1 = Small, 2 = Medium, 3 = Large)")
	fmt.Scanln(&sizeChoice)

	var drinkChoice string
	fmt.Println("\nWhat would you like to drink? (leave blank to keep Coke as default)")
	fmt.Scanln(&drinkChoice)

	c.SetSize(sizeChoice)
	c.SetDrink(drinkChoice)
	return c
}

func main() {
	fmt.Println("Welcome to Eclectic Drive-Thru. What would you like to order?")
	var selectedCombo Combo = PromptMenuAndSelectCombo()
	selectedCombo = UpdateSizeAndDrink(selectedCombo)

	fmt.Println("\nHere is your order:")
	selectedCombo.DisplayFull()
	fmt.Println("\n\nThank you, please pull forward")
}
