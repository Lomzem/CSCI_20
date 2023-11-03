/***************************************
Joust
Author: Lawjay Lee
Date Completed: 2023-11-02
Description: Simulate a jousting match between two knights
***************************************/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

type Knight struct {
	Name    string
	Stamina int
	Weapon  Weapon
	Mounted bool
}

type Weapon struct {
	StaminaCost int
	HitChance   int
	Type        string
}

func NewKnight(name string, stamina int, weapon Weapon) Knight {
	var k Knight
	k.Name = name
	k.Stamina = stamina
	k.Weapon = weapon
	k.Mounted = true
	return k
}

func NewWeapon(staminaCost int, hitChance int, _type string) Weapon {
	var w Weapon
	w.StaminaCost = staminaCost
	w.HitChance = hitChance
	w.Type = _type
	return w
}

func (w Weapon) GetHitChance() int {
	return w.HitChance
}
func (w Weapon) GetStaminaCost() int {
	return w.StaminaCost
}

func (w Weapon) GetType() string {
	return w.Type
}

func (w *Weapon) SetHitChance(chance int) {
	w.HitChance = chance
}

func (w *Weapon) SetStaminaCost(cost int) {
	w.StaminaCost = cost
}

func (w *Weapon) SetType(name string) {
	w.Type = name
}

func (k Knight) GetMounted() bool {
	return k.Mounted
}

func (k Knight) GetName() string {
	return k.Name
}

func (k Knight) GetStamina() int {
	return k.Stamina
}

func (k Knight) GetWeapon() Weapon {
	return k.Weapon
}

func (k *Knight) SetMounted(status bool) {
	k.Mounted = status
}

func (k *Knight) SetName(name string) {
	k.Name = name
}

func (k *Knight) SetStamina(newStamina int) {
	k.Stamina = newStamina
}

func (k *Knight) SetWeapon(weapon Weapon) {
	k.Weapon = weapon
}

func (k Knight) isExhausted() bool {
	return k.GetStamina() <= 0
}

func (k Knight) DisplayStats() {
	if !k.isExhausted() && k.GetMounted() {
		fmt.Printf("%s is not exhausted (stamina=%d) and is still on horse.\n", k.GetName(), k.GetStamina())
	} else if k.isExhausted() && k.GetMounted() {
		fmt.Printf("%s is exhausted and is still on horse.\n", k.GetName())
	} else if !k.isExhausted() && !k.GetMounted() {
		fmt.Printf("%s is not exhausted (stamina=%d) and has been knocked off of the horse.\n", k.GetName(), k.GetStamina())
	} else if k.isExhausted() && !k.GetMounted() {
		fmt.Printf("%s is exhausted and has been knocked off of the horse.\n", k.GetName())
	}

	fmt.Printf("%s is using: %s that requires %d stamina and has a %d%% chance to hit.\n", k.GetName(), k.Weapon.GetType(), k.Weapon.GetStaminaCost(), k.Weapon.GetHitChance())
}

func (k *Knight) Joust() bool {
	var newStamina int = k.GetStamina() - k.Weapon.GetStaminaCost()
	k.SetStamina(newStamina)

	var successfulHit bool = k.Weapon.Swing()
	return successfulHit
}

func (w Weapon) Swing() bool {
	var hitChance int = w.GetHitChance()
	var successfulHit bool = hitChance >= rand.Intn(100)+1
	return successfulHit
}

func main() {
	var Knight1Name string
	var Knight1Stamina int
	var Knight1WeaponType string
	var Knight1HitChance int
	var Knight1StaminaCost int

	var Knight2Name string
	var Knight2Stamina int
	var Knight2WeaponType string
	var Knight2HitChance int
	var Knight2StaminaCost int

	flag.StringVar(&Knight1Name, "Knight1Name", "", "")
	flag.IntVar(&Knight1Stamina, "Knight1Stamina", 0, "")
	flag.StringVar(&Knight1WeaponType, "Knight1WeaponType", "", "")
	flag.IntVar(&Knight1HitChance, "Knight1HitChance", 0, "")
	flag.IntVar(&Knight1StaminaCost, "Knight1StaminaCost", 0, "")

	flag.StringVar(&Knight2Name, "Knight2Name", "", "")
	flag.IntVar(&Knight2Stamina, "Knight2Stamina", 0, "")
	flag.StringVar(&Knight2WeaponType, "Knight2WeaponType", "", "")
	flag.IntVar(&Knight2HitChance, "Knight2HitChance", 0, "")
	flag.IntVar(&Knight2StaminaCost, "Knight2StaminaCost", 0, "")

	flag.Parse()

	var weapon1 Weapon = NewWeapon(Knight1StaminaCost, Knight1HitChance, Knight1WeaponType)
	var knight1 Knight = NewKnight(Knight1Name, Knight1Stamina, weapon1)

	var weapon2 Weapon = NewWeapon(Knight2StaminaCost, Knight2HitChance, Knight2WeaponType)
	var knight2 Knight = NewKnight(Knight2Name, Knight2Stamina, weapon2)

	var knight1Hit bool
	var knight2Hit bool

	for knight1.GetMounted() && knight2.GetMounted() {
		knight1Hit = knight1.Joust()
		knight2.SetMounted(!knight1Hit)

		knight2Hit = knight2.Joust()
		knight1.SetMounted(!knight2Hit)

		knight1.DisplayStats()
		knight2.DisplayStats()
		fmt.Println()

		if knight1.isExhausted() {
			// end the loop if the knight is exhausted, too tired to fight so gets off horse
			// should be fine to SetMounted here because the correct Mounted status is already
			// printed and won't call DisplayStats with incorrect info next iteration because
			// loop will break since knight.GetMounted() should no longer return true
			knight1.SetMounted(false)
		}

		if knight2.isExhausted() {
			knight2.SetMounted(false)
		}
	}

	// check if knights are in a "losing state", either have no stamina or off their horse
	// being exhausted changed mount state of knights, so shouldn't have to check for it here
	if !knight1.GetMounted() && !knight2.GetMounted() {
		fmt.Println("It's a draw!")
	} else if !knight2.GetMounted() {
		fmt.Println(knight1.GetName(), "wins!")
	} else if !knight1.GetMounted() {
		fmt.Println(knight2.GetName(), "wins!")
	}
}
