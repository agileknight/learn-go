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

type CamelStartPositioner interface {
	Position(camels []CamelState)
}

type RandomCamelStartPositioner struct {
	camelStepDice Dice
}

func (this *RandomCamelStartPositioner) Position(camels []CamelState) {
	camelsByPosition := make(map[int][]CamelState)
	for i := range camels {
		camel := &camels[i]
		pos := this.camelStepDice.Roll()
		camel.position = pos

		maxLevel := -1
		for _, camelAtPos := range camelsByPosition[pos] {
			if camelAtPos.level > maxLevel {
				maxLevel = camelAtPos.level
			}
		}

		camel.level = maxLevel + 1
		camelsByPosition[pos] = append(camelsByPosition[pos], *camel)
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

func (this *Game) Dice() {
	// TODO implement
}
