package game

type Strategy interface {
	RunStrategy() bool
}

//To check if all slots are equal
type threeSlots struct {
}

func (t *threeSlots) RunStrategy(p *Player) bool {
	if len(p.History) == 1 {
		return true
	}
	return false
}

//3 times two same (449 994 885)
type twoSlots struct {
}

func (t *twoSlots) RunStrategy(p *Player) bool {
	if len(p.History) < JackpotInstance.GetChances() {
		return true
	}
	return false
}
