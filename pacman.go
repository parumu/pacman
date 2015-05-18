package main

type pacman struct {
	loc        xy
	vec        xy
	extraPower int
}

func (p *pacman) move(loc xy, vec xy) {
	p.loc = loc
	p.vec = vec

	if p.extraPower > 0 {
		p.extraPower--
	}
}

func (p *pacman) hasExtraPower() bool {
	return p.extraPower > 0
}

func (p *pacman) addExtraPower(i int) {
	p.extraPower += i
}

func buildPacman() pacman {
	return pacman{xy{}, vecLeft, 0}
}
