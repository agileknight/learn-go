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
	camelMovesLeft int
}

type CamelStartPositioner interface {
	Position(camels []CamelState)
}

type RandomCamelStartPositioner struct {
	camelStepDice Dice
}

func findInsertLevelAtPos(camels []CamelState, position int) int {
	maxLevel := -1
	for _, camel := range camels {
		if camel.position == position && camel.level > maxLevel {
			maxLevel = camel.level
		}
	}
	return maxLevel + 1
}

func (this *RandomCamelStartPositioner) Position(camels []CamelState) {
	for i := range camels {
		pos := this.camelStepDice.Roll()
		camels[i].level = findInsertLevelAtPos(camels, pos)
		camels[i].position = pos
	}
}

type Game struct {
	config               GameConfig
	camelIndexDice       Dice
	camelStepDice        Dice
	camelStartPositioner CamelStartPositioner
	state                GameState
}

func Init(config GameConfig) *Game {
	camelStepDice := BoundedDice{&RandomRandInt{}, config.minCamelSteps, config.maxCamelSteps}
	game := Game{
		config: config,

		camelIndexDice:       &NoDuplicatesBoundedDice{&RandomRandInt{}, 0, config.numCamels, nil},
		camelStepDice:        &camelStepDice,
		camelStartPositioner: &RandomCamelStartPositioner{&camelStepDice},

		state: GameState{
			players:        make([]PlayerState, config.numPlayers),
			camels:         make([]CamelState, config.numCamels),
			curPlayerIndex: config.startPlayerIndex,
			camelMovesLeft: config.numCamels,
		},
	}
	for i := range game.state.players {
		game.state.players[i].money = config.playerStartMoney
	}
	game.camelStartPositioner.Position(game.state.camels)
	return &game
}

func (this *Game) Bet(camelIndex int) {
	// TODO implement
}

func moveCamel(camels []CamelState, camelIndex int, camelSteps int) {
	// TODO implement
}

func (this *Game) Dice() {
	camelIndex := this.camelIndexDice.Roll()
	camelSteps := this.camelStepDice.Roll()
	moveCamel(this.state.camels, camelIndex, camelSteps)
}
