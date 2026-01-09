package task1

type PepeSchenele struct {
	Speed    int
	Charisma int
	Wisdom   int
}

func NewPepeSchenele(speed int, charisma int, wisdom int) *PepeSchenele {
	return &PepeSchenele{
		Speed:    speed,
		Charisma: charisma,
		Wisdom:   wisdom,
	}
}

func (p *PepeSchenele) GetRaiting() int {
	return (p.Speed * 2) + (p.Charisma * 3) + p.Wisdom
}
