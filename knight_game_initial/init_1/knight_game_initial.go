package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Weapon struct {
	Name       string
	hitChance  int
	staminaReq int
}

type Knight struct {
	Name    string
	Stamina int
	Lance   Weapon
	Win     bool
}

func (k *Knight) getStaminaStr() string {
	return "(stamina=" + strconv.Itoa(k.Stamina) + ")"
}

func (w *Weapon) weaponStatStr() string {
	var hitChanceStr string = strconv.Itoa(w.hitChance)
	return w.Name + " that requires " + strconv.Itoa(w.staminaReq) + " stamina and has a " + hitChanceStr + "% chance to hit."
}

func (k *Knight) depleteStamina() {
	k.Stamina -= k.Lance.staminaReq
}

func (k *Knight) isExhausted() bool {
	return k.Stamina <= 0
}

func (k *Knight) attack() bool {
	return k.Lance.hitChance >= rand.Intn(100)+1
}

func (k *Knight) playRound() bool {
	k.depleteStamina()
	switch k.isExhausted() {
	case true:
		k.DisplayExhausted()
		return false
	case false:
		k.DisplayNotExhausted()
		var successfulHit bool = k.attack()
		if successfulHit {
			fmt.Println(k.Name, "got a successful hit!")
			k.Win = true
		}
	}
	return true
}

func (k *Knight) DisplayNotExhausted() {
	fmt.Println(k.Name, "is not exhausted "+k.getStaminaStr()+" and is still on the horse.")
	fmt.Println(k.Name, "is using:", k.Lance.weaponStatStr())
}

func (k *Knight) DisplayExhausted() {
	fmt.Println("is exhausted and lost")
}

func (k *Knight) WinMsg() {
	fmt.Println(k.Name, "has won!")
}

func main() {
	var lance Weapon = Weapon{"Lance", 50, 3}
	var sword Weapon = Weapon{"Sword", 50, 2}
	var knight1 Knight = Knight{"Sir Cuss", 20, lance, false}
	var knight2 Knight = Knight{"Lancelot", 20, sword, false}

	for !knight1.Win && !knight2.Win {
		knight1.playRound()
		fmt.Println()
		knight2.playRound()
	}
	if knight1.Win && knight2.Win {
		fmt.Println("It's a draw")
	} else if knight1.Win {
		knight1.WinMsg()
	} else {
		knight2.WinMsg()
	}
}
