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
		config := GameConfig{
			playerStartMoney: 3,
			numPlayers:       c.numPlayers,
		}

		game := Init(config)
		if spliceLen := len(game.players); spliceLen != c.numPlayers {
			t.Errorf("Player splice was size %d for num players %d", spliceLen, c.numPlayers)
		}

		for i, p := range game.players {
			if p.money != config.playerStartMoney {
				t.Errorf("Player %d got initialized with money %d not matching start amount %d", i, p.money, config.playerStartMoney)
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
