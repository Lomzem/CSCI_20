/***************************************
Concepts Review
Author: Lawjay Lee
Date Completed: 2023-10-04
Description: Creates three pets then displays their info
***************************************/

package main

import "fmt"

type Pet struct {
	Name      string
	OwnerName string
	Type      string
	Age       int
}

func NewPet(name string, owner string, type_ string, age int) Pet {
	var newPet Pet
	newPet.Name = name
	newPet.OwnerName = owner
	newPet.Type = type_
	newPet.Age = age
	return newPet
}

func (p *Pet) GetName() string {
	return p.Name
}

func (p *Pet) SetName(name string) {
	p.Name = name
}

func (p *Pet) GetOwnerName() string {
	return p.OwnerName
}

func (p *Pet) SetOwnerName(name string) {
	p.OwnerName = name
}

func (p *Pet) GetType() string {
	return p.Type
}

func (p *Pet) SetType(type_ string) {
	p.Type = type_
}

func (p *Pet) GetAge() int {
	return p.Age
}

func (p *Pet) SetAge(age int) {
	p.Age = age
}

func (p *Pet) Display(petNum int) {
	fmt.Printf("Pet %d:\n", petNum)
	fmt.Println("Name:", p.GetName())
	fmt.Println("Owner:", p.GetOwnerName())
	fmt.Println("Type:", p.GetType())
	fmt.Println("Age:", p.GetAge())
	fmt.Println()
}

func BuildPetFromUserInput() Pet {
	var name string
	fmt.Println("Enter pet name")
	fmt.Scanln(&name)

	var owner string
	fmt.Println("Enter pet owner name")
	fmt.Scanln(&owner)

	var type_ string
	fmt.Println("Enter pet type")
	fmt.Scanln(&type_)

	var age int
	fmt.Println("Enter pet age")
	fmt.Scanln(&age)
	fmt.Println()

	var newPet Pet = NewPet(name, owner, type_, age)
	return newPet
}

func main() {
	var pet1 Pet = BuildPetFromUserInput()
	var pet2 Pet = BuildPetFromUserInput()
	var pet3 Pet = BuildPetFromUserInput()

	pet1.Display(1)
	pet2.Display(2)
	pet3.Display(3)

	if pet1.GetAge() >= 10 {
		fmt.Println(pet1.GetName(), "is old")
	} else {
		fmt.Println(pet1.GetName(), "is young")
	}

	if pet2.GetAge() >= 10 {
		fmt.Println(pet2.GetName(), "is old")
	} else {
		fmt.Println(pet2.GetName(), "is young")
	}

	if pet3.GetAge() >= 10 {
		fmt.Println(pet3.GetName(), "is old")
	} else {
		fmt.Println(pet3.GetName(), "is young")
	}
}
