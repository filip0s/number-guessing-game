package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	player1         Player
	player2         Player
	generatedNumber int
}

// Provides initial values for Game struct and runs the game
func (g *Game) create() {
	g.player1.create("Player 1")
	g.player2.create("Player 2")

	rand.Seed(time.Now().UnixNano())
	g.generatedNumber = rand.Intn(MaximalGuessValue)

	g.run()
}

func (g *Game) collectPlayerInputs() {
	g.player1.readInput()
	g.player2.readInput()
}

// Calculates difference between guess provided by Player struct in parameter and
// generated guess in the Game struct. Because absolute value is expected
// returned value of this function, negative differences are multiplied by -1.
func (g *Game) calculateAbsoluteDifference(player Player) int {
	var diff int
	diff = player.guess - g.generatedNumber

	if diff < 0 {
		return -1 * diff
	}

	return diff
}

// Compares which of two Game players was closer to the guess
func (g *Game) decideWinner() {
	var (
		p1diff   = g.calculateAbsoluteDifference(g.player1)
		p2diff   = g.calculateAbsoluteDifference(g.player2)
		decision = "It's a tie!"
	)

	if p1diff < p2diff {
		decision = fmt.Sprintf("%s won", g.player1.name)
	}

	if p2diff < p1diff {
		decision = fmt.Sprintf("%s won", g.player2.name)
	}

	fmt.Println(decision)
}

func (g *Game) run() {
	g.collectPlayerInputs()
	g.decideWinner()

}
