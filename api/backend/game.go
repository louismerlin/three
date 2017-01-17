package backend

//import

type Player struct {
	client    *Client
	lives     int
	hasPlayed bool
}

type Game struct {
	players [3]*Player
	rows    map[int]int
}

func NewGame(clients [3]*Client) *Game {
  return &Game{
    [3]*Player{
      &Player{ clients[0], 3, false },
      &Player{ clients[1], 3, false },
      &Player{ clients[2], 3, false },
    },
    make(map[int]int)
  }
}

func (*Game) Receive(msg *Message) {

}
