package player

type player struct {
	Name     string
	Attempts int
	History  []Outcome
}

type Outcome struct {
	AttemptNumber int         //indexes
	Result        map[int]int //frequency map
}

type Player interface {
	PullLever()
}

func NewPlayer(name string) Player {
	return &player{
		Name:     name,
		Attempts: 0,
	}
}

func (p *player) PullLever() {
	p.Attempts++
	//
}
