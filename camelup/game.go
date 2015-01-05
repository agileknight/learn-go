package camelup

import (
	"math/rand"
)

const (
	PlayerStartMoney int = 3
	NumCamels int = 2
	NumPlayers int = 2
	MinCamelSteps int = 1
	MaxCamelSteps int = 3
)

type Player struct {
	money int
}

// Returns a random int in [0..n)
type RandInt interface {
	Intn(n int) int
}

type RandomRandInt struct {
	
}

func (this *RandomRandInt) Intn(n int) int {
	return rand.Intn(n)
}

type Dice interface {
	Roll() int
}

type BoundedDice struct {
	rand RandInt
	minValue int
	maxValue int
}

func (this *BoundedDice) Roll() int {
	return this.minValue + this.rand.Intn(this.maxValue - this.minValue + 1)
}

type Game struct {
	players []Player
	camelIndexDice Dice
	camelStepDice Dice
}

func Init(numPlayers int) *Game {
	game := Game{
		players: make([]Player, numPlayers),
	}
	for i := range game.players {
		game.players[i].money = PlayerStartMoney
	}
	return &game
}

func (this *Game) Bet(camelIndex int) {
	// TODO implement
}

func (this *Game) Dice() {
	// TODO implement
}
