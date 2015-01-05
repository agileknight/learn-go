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

type CamelDice interface {
	Roll() (camelIndex, steps int)
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

type RandomCamelDice struct {
	randInt RandInt
	numCamels int
	minSteps int
	maxSteps int
}

func (this *RandomCamelDice) Roll() (camelIndex, steps int) {
	return this.randInt.Intn(this.numCamels), this.randInt.Intn(this.maxSteps)
}

type Game struct {
	players []Player
	camelDice CamelDice
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
