package camelup

const (
	PlayerStartMoney int = 3
	NumCamels        int = 2
	NumPlayers       int = 2
	MinCamelSteps    int = 1
	MaxCamelSteps    int = 3
)

type Player struct {
	money int
}

type Dice interface {
	Roll() int
}

type Game struct {
	players        []Player
	camelIndexDice Dice
	camelStepDice  Dice
}

func Init(numPlayers int) *Game {
	game := Game{
		players: make([]Player, numPlayers),
	}
	for i := range game.players {
		game.players[i].money = PlayerStartMoney
	}
	return &game
}

func (this *Game) Bet(camelIndex int) {
	// TODO implement
}

func (this *Game) Dice() {
	// TODO implement
}
