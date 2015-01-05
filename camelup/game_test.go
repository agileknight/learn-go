package camelup

import "testing"

func TestInit(t *testing.T) {
	cases := []struct {
		numPlayers int
	}{
		{2},
		{4},
	}

	for _, c := range cases {
		game := Init(c.numPlayers)
		if spliceLen := len(game.players); spliceLen != c.numPlayers {
			t.Errorf("Player splice was size %d for num players %d", spliceLen, c.numPlayers)
		}

		for i, p := range game.players {
			if p.money != PlayerStartMoney {
				t.Errorf("Player %d got initialized with money %d not matching start amount %d", i, p.money, PlayerStartMoney)
			}
		}
	}
}

type MockRandInt struct {
	returnVal int
	calledWith int
}

func (this *MockRandInt) Intn(n int) int {
	this.calledWith = n
	return this.returnVal
}

func (this *MockRandInt) Returns(n int) {
	this.returnVal = n
}

func (this *MockRandInt) CalledWith() int {
	return this.calledWith
} 

func TestBoundedDiceRoll(t *testing.T) {
	rand := MockRandInt{}
	dice := BoundedDice{
		rand: &rand,
		minValue: 3,
		maxValue: 5,
	}
	
	cases := []struct{
		randResult int
		expectedParam int
		expectedResult int
	}{
		{0, 3, 3},
		{1, 3, 4},
		{2, 3, 5},
	}
	
	for _, c := range cases {
		rand.Returns(c.randResult)
		got := dice.Roll()
		if got != c.expectedResult {
			t.Errorf("Expected %d but found %d", c.expectedResult,  got)
		}
		if param := rand.CalledWith(); param != c.expectedParam {
			t.Errorf("Expected param %d but found %d", c.expectedParam, param)
		}
	}
}

func TestAccFullBettingRound(t *testing.T) {
	// TODO setup with 2 camels in known starting position
	game := Init(2)
	game.Bet(0)
	// TODO setup so that camel 0 gets ahead
	game.Dice()
	game.Dice()
	if game.players[0].money != 8 {
		t.Errorf("Bet was not counted.")
	}
}
