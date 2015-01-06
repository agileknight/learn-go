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
	money                  int
	betAmountsByCamelIndex map[int][]int
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
		game.state.players[i].betAmountsByCamelIndex = make(map[int][]int)
	}
	game.camelStartPositioner.Position(game.state.camels)
	return &game
}

func findCamelAtPosAndLevel(camels []CamelState, pos int, level int) (index int, found bool) {
	for i := range camels {
		if camels[i].position == pos && camels[i].level == level {
			return i, true
		}
	}
	return -1, false
}

func moveCamelsStartingAtLevel(camels []CamelState, curPos int, targetPos int, level int) {
	i, found := findCamelAtPosAndLevel(camels, curPos, level)
	if found {
		targetLevel := findInsertLevelAtPos(camels, targetPos)
		camels[i].level = targetLevel
		camels[i].position = targetPos
		moveCamelsStartingAtLevel(camels, curPos, targetPos, level+1)
	}
}

func moveCamel(camels []CamelState, camelIndex int, camelSteps int) {
	moveLevel := camels[camelIndex].level
	curPos := camels[camelIndex].position
	targetPos := curPos + camelSteps
	moveCamelsStartingAtLevel(camels, curPos, targetPos, moveLevel)
}

func findWinnerExcludingIndex(camels []CamelState, excludedIndex int) int {
	winnerIndex := 0
	winPos := 0
	winLevel := 0
	for i := range camels {
		if i == excludedIndex {
			continue
		}
		if camels[i].position > winPos || camels[i].position == winPos && camels[i].level > winLevel {
			winPos = camels[i].position
			winLevel = camels[i].level
			winnerIndex = i
		}
	}

	return winnerIndex
}

func findWinnerCamels(camels []CamelState) (winnerIndex int, secondIndex int) {
	winnerIndex = findWinnerExcludingIndex(camels, -1)
	secondIndex = findWinnerExcludingIndex(camels, winnerIndex)
	return
}

func (this *Game) payout() {
	winnerIndex, secondIndex := findWinnerCamels(this.state.camels)
	for _, player := range this.state.players {
		totalDiff := 0
		for _, amount := range player.betAmountsByCamelIndex[winnerIndex] {
			totalDiff += amount
		}
		totalDiff += len(player.betAmountsByCamelIndex[secondIndex])
		for i := range this.state.camels {
			if i != winnerIndex && i != secondIndex {
				totalDiff -= len(player.betAmountsByCamelIndex[i])
			}
		}
	}
}

func (this *Game) nextPlayer() {
	this.state.curPlayerIndex = (this.state.curPlayerIndex + 1) % this.config.numPlayers
}

func (this *Game) Dice() {
	camelIndex := this.camelIndexDice.Roll()
	camelSteps := this.camelStepDice.Roll()
	moveCamel(this.state.camels, camelIndex, camelSteps)
	this.state.camelMovesLeft--
	if this.state.camelMovesLeft == 0 {
		this.state.camelMovesLeft = this.config.numCamels
		this.payout()
	}
	this.nextPlayer()
}

func (this *Game) Bet(camelIndex int) {
	this.state.players[this.state.curPlayerIndex].betAmountsByCamelIndex[camelIndex] = append(this.state.players[this.state.curPlayerIndex].betAmountsByCamelIndex[camelIndex], 5)
	this.nextPlayer()
}
