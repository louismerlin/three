package backend

type Message struct {
	C1 int `json:"choice1"`
	C2 int `json:"choice2"`
	C3 int `json:"choice3"`
}

func (self *Message) String() string {
	return "Choices received !"
}
