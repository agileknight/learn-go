package camelup

import "testing"

type StubDice struct {
	rolls   []int
	curRoll int
}

func (this *StubDice) Roll() int {
	result := this.rolls[this.curRoll]
	this.curRoll++
	return result
}

func TestAccFullBettingRound(t *testing.T) {
	config := GameConfig{
		playerStartMoney: 3,
		numCamels:        2,
		numPlayers:       2,
		minCamelSteps:    1,
		maxCamelSteps:    3,
		boardLength:      20,
		startPlayerIndex: 0,
	}

	game := Init(config)

	game.camelIndexDice = &StubDice{rolls: []int{0, 1}}
	game.camelStepDice = &StubDice{rolls: []int{3, 1}}

	game.Bet(0)
	game.Dice()
	game.Dice()
	if game.state.players[0].money != 8 {
		t.Errorf("Bet was not counted.")
	}
}
