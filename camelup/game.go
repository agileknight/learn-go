package camelup

type GameConfig struct {
	playerStartMoney int
	numCamels        int
	numPlayers       int
	minCamelSteps    int
	maxCamelSteps    int
}

type Player struct {
	money int
}

type Dice interface {
	Roll() int
}

type Game struct {
	config         GameConfig
	players        []Player
	camelIndexDice Dice
	camelStepDice  Dice
}

func Init(config GameConfig) *Game {
	game := Game{
		config:  config,
		players: make([]Player, config.numPlayers),
	}
	for i := range game.players {
		game.players[i].money = config.playerStartMoney
	}
	return &game
}

func (this *Game) Bet(camelIndex int) {
	// TODO implement
}

func (this *Game) Dice() {
	// TODO implement
}
