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

type PlayerState struct {
	money int
}

// position is the index on the board
// level is the height (for stacking camels)
type CamelState struct {
	position int
	level    int
}

type Dice interface {
	Roll() int
}

type GameState struct {
	players        []PlayerState
	camels         []CamelState
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
			players:        make([]PlayerState, config.numPlayers),
			camels:         make([]CamelState, config.numCamels),
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
