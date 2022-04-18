package game

type slot struct {
	ID int
}

type jackpot struct {
	Slots   []slot
	Chances int
	//players history
}

type Jackpot interface {
	SpinTheWheel() []slot
	GetChances() int
}

//singleton class
func NewJackpot(slots int, chances int) Jackpot {
	return &jackpot{
		Slots:   make([]slot, slots),
		Chances: chances,
	}
}
