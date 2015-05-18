package main

type pacman struct {
	loc        xy
	vec        vec
	extraPower int
}

func (p *pacman) liveLife() {
	if p.extraPower > 0 {
		p.extraPower--
	}
}

func (p *pacman) move(loc xy, vec vec) {
	p.loc = loc
	p.vec = vec
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
