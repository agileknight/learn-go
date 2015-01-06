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

type NoDuplicatesBoundedDice struct {
	rand        RandInt
	minValue    int
	maxValue    int
	resultsLeft []int
}

func (this *NoDuplicatesBoundedDice) Roll() int {
	if len(this.resultsLeft) == 0 {
		numValues := this.maxValue - this.minValue + 1
		this.resultsLeft = make([]int, numValues)
		for i := range this.resultsLeft {
			this.resultsLeft[i] = this.minValue + i
		}
	}
	
	numLeft := len(this.resultsLeft)
	resultIndex := this.rand.Intn(numLeft)
	result := this.resultsLeft[resultIndex]
	this.resultsLeft = append(this.resultsLeft[:resultIndex], this.resultsLeft[resultIndex+1:]...)
	return result
}
