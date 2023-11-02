/***************************************
Joust
Author: Lawjay Lee
Date Completed: 2023-11-02
Description: Simulate a jousting match between two knights
***************************************/

package main

import "fmt"

type Knight struct {
	Name       string
	Stamina    int
	WeaponType Weapon
	Mounted    bool
}

type Weapon struct {
	Name        string
	StaminaCost int
	HitChance   int
}

func NewKnight(name string, stamina int, weaponType Weapon) Knight {
	var k Knight
	k.Name = name
	k.Stamina = stamina
	k.WeaponType = weaponType
	k.Mounted = true
	return k
}

func NewWeapon(name string, staminaCost int, hitChance int) Weapon {
	var w Weapon
	w.Name = name
	w.StaminaCost = staminaCost
	w.HitChance = hitChance
	return w
}

func (k *Knight) SetStamina(newStamina int) {
	k.Stamina = newStamina
}

func (k *Knight) isMounted() bool {
	return k.Mounted
}

func (k *Knight) isExhausted() bool {
	return k.Stamina <= 0
}

func (k *Knight) PrintStats() {
	var exhaustedString string
	switch k.isExhausted() {
	case true:
		exhaustedString = "exhausted"
	case false:
		exhaustedString = "not exhausted"
	}

	switch k.isMounted() {
	case true:
		fmt.Printf("%s is %s (stamina=%d) and is still on the horse.\n", k.Name, exhaustedString, k.Stamina)
		fmt.Printf("%s is using: %s that requires %d stamina and has a %d%% chance to hit.\n\n", k.Name, k.WeaponType.Name, k.WeaponType.StaminaCost, k.WeaponType.HitChance)
	}
}

func (k *Knight) Swing() {
}

func main() {
	Knight1Name := "King Arthur"
	Knight1Stamina := 50
	Knight1WeaponType := "Excalibur"
	Knight1HitChance := 15
	Knight1StaminaCost := 10
	Knight2Name := "Black Knight"
	Knight2Stamina := 40
	Knight2WeaponType := "Longsword"
	Knight2HitChance := 10
	Knight2StaminaCost := 5

	var weapon1 Weapon = NewWeapon(Knight1WeaponType, Knight1StaminaCost, Knight1HitChance)
	var knight1 Knight = NewKnight(Knight1Name, Knight1Stamina, weapon1)

	var weapon2 Weapon = NewWeapon(Knight2WeaponType, Knight2StaminaCost, Knight2HitChance)
	var knight2 Knight = NewKnight(Knight2Name, Knight2Stamina, weapon2)
	for knight1.isMounted() && knight2.isMounted() {
		knight1.PrintStats()
		knight2.PrintStats()
		knight1.Mounted = false
	}
}
