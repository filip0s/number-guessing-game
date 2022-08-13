package main

import (
	"fmt"
	"os"
)

type Player struct {
	name  string
	guess int
}

// Provides initial values for the Player struct
func (p *Player) create(name string) {
	p.name = name
}

func (p *Player) readInput() {
	fmt.Printf("%s, enter your guess: ", p.name)
	var input int
	_, err := fmt.Scanf("%d", &input)

	// Guard clause to exit early if there are no errors from reading user input
	if err == nil {
		p.guess = input
		return
	}

	_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)

	p.readInput()
}
