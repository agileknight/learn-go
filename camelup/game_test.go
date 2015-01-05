package camelup

import "testing"

func TestInit(t *testing.T) {
	cases := []struct {
		numPlayers int
	}{
		{2},
		{4},
	}
	
	for _,c := range cases {
		game := Game{}
		game.Init(c.numPlayers)
		if spliceLen := len(game.players); spliceLen != c.numPlayers {
			t.Errorf("Player splice was size %d for num players %d", spliceLen, c.numPlayers)
		}
	}
}