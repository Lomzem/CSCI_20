/***************************************
Dice Game
Author: Lawjay Lee
Date Completed: 2023-10-18
Description: Game where 2 players roll 2 six-sided die
***************************************/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Implement Die data type definition, related member functions, and pseudo-constructor here
type Die struct {
	Sides int
}

func NewDie(sides int) Die {
	var d Die
	d.Sides = sides
	return d
}

func (d *Die) Roll() int {
	return rand.Intn(d.Sides) + 1
}

// Implement Player data type definition, related member functions, and pseudo-constructor here
type Player struct {
	Die1 Die
	Die2 Die
}

func NewPlayer() Player {
	var p Player
	p.Die1 = NewDie(6)
	p.Die2 = NewDie(6)
	return p
}

func (p *Player) RollDice() int {
	return p.Die1.Roll() + p.Die2.Roll()
}

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func main() {

	var player1 Player = NewPlayer()
	var player2 Player = NewPlayer()

	var p1Score int = player1.RollDice()
	var p2Score int = player2.RollDice()

	// Implement remaining game logic here
	fmt.Println("Player 1 Roll:", p1Score)
	fmt.Println("Player 2 Roll:", p2Score)

	if p1Score > p2Score {
		fmt.Println("Player 1 Wins!")
	} else if p2Score > p1Score {
		fmt.Println("Player 2 Wins!")
	} else {
		fmt.Println("It's a tie!")
	}
}
