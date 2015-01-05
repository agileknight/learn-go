package camelup

type Player struct {
	money int
}

type Game struct {
	players []Player
}

func (this *Game) Init(numPlayers int) {
	this.players = make([]Player, numPlayers)
}