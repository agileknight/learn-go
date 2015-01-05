package camelup

type GameConfig struct {
	playerStartMoney int
	numCamels        int
	numPlayers       int
	minCamelSteps    int
	maxCamelSteps    int
	boardLength      int
	startPlayerIndex int
}

type Player struct {
	money int
}

type Dice interface {
	Roll() int
}

type GameState struct {
	players        []Player
	camelPositions []int
	camelOrder     []int
	curPlayerIndex int
}

type Game struct {
	config         GameConfig
	camelIndexDice Dice
	camelStepDice  Dice
	state          GameState
}

func Init(config GameConfig) *Game {
	game := Game{
		config: config,

		camelIndexDice: &BoundedDice{&RandomRandInt{}, 0, config.numCamels},
		camelStepDice:  &BoundedDice{&RandomRandInt{}, config.minCamelSteps, config.maxCamelSteps},

		state: GameState{
			players:        make([]Player, config.numPlayers),
			camelPositions: make([]int, config.numCamels),
			camelOrder:     make([]int, config.numCamels),
			curPlayerIndex: config.startPlayerIndex,
		},
	}
	for i := range game.state.players {
		game.state.players[i].money = config.playerStartMoney
	}
	return &game
}

func (this *Game) Bet(camelIndex int) {
	// TODO implement
}

func (this *Game) Dice() {
	// TODO implement
}
