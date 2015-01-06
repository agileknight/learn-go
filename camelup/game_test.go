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

func TestStartPositioning(t *testing.T) {
	cases := []struct {
		numCamels         int
		diceRolls         []int
		expectedPositions []int
		expectedLevels    []int
	}{
		{3, []int{1, 3, 2}, []int{1, 3, 2}, []int{0, 0, 0}},
		{2, []int{1, 1}, []int{1, 1}, []int{0, 1}},
	}

	for _, c := range cases {
		positioner := &RandomCamelStartPositioner{camelStepDice: &StubDice{rolls: c.diceRolls}}
		camelStates := make([]CamelState, c.numCamels)
		positioner.Position(camelStates)
		for i := range camelStates {
			if gotPos := camelStates[i].position; gotPos != c.expectedPositions[i] {
				t.Errorf("Wrong starting position of camel at index %d. Expected %d but was %d.", i, c.expectedPositions[i], gotPos)
			}
			if gotLevel := camelStates[i].level; gotLevel != c.expectedLevels[i] {
				t.Errorf("Wrong starting level of camel at index %d. Expected %d but was %d.", i, c.expectedLevels[i], gotLevel)
			}
		}
	}
}

func TestMoveCamel(t *testing.T) {
	cases := []struct {
		camelPositions    []int
		camelLevels       []int
		index             int
		steps             int
		expectedPositions []int
		expectedLevels    []int
	}{
		{[]int{0, 1}, []int{0, 0}, 0, 1, []int{1, 1}, []int{1, 0}},
		{[]int{0, 0}, []int{0, 1}, 0, 1, []int{1, 1}, []int{0, 1}},
	}

	for _, c := range cases {
		numCamels := len(c.camelPositions)
		camelStates := make([]CamelState, numCamels)
		for i := range camelStates {
			camelStates[i] = CamelState{c.camelPositions[i], c.camelLevels[i]}
		}
		moveCamel(camelStates, c.index, c.steps)
		for i := range camelStates {
			if gotPos := camelStates[i].position; gotPos != c.expectedPositions[i] {
				t.Errorf("Position mismatch of camel at index %d. Expected %d but was %d.", i, c.expectedPositions[i], gotPos)
			}
			if gotLevel := camelStates[i].level; gotLevel != c.expectedLevels[i] {
				t.Errorf("Level mismatch of camel at index %d. Expected %d but was %d.", i, c.expectedLevels[i], gotLevel)
			}
		}
	}
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
	game.camelStartPositioner = &RandomCamelStartPositioner{camelStepDice: &StubDice{rolls: []int{0, 1}}}

	game.Bet(0)
	game.Dice()
	game.Dice()
	if game.state.players[0].money != 8 {
		t.Errorf("Bet was not counted.")
	}
}
