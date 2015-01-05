package camelup

import "testing"

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
	game.Bet(0)
	// TODO setup so that camel 0 gets ahead
	game.Dice()
	game.Dice()
	if game.players[0].money != 8 {
		t.Errorf("Bet was not counted.")
	}
}
