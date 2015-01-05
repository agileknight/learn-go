package camelup

import (
	"math/rand"
)

// Returns a random int in [0..n)
type RandInt interface {
	Intn(n int) int
}

type RandomRandInt struct {
}

func (this *RandomRandInt) Intn(n int) int {
	return rand.Intn(n)
}

type BoundedDice struct {
	rand     RandInt
	minValue int
	maxValue int
}

func (this *BoundedDice) Roll() int {
	return this.minValue + this.rand.Intn(this.maxValue-this.minValue+1)
}
