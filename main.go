package main

import "time"

var vecLeft, vecRight, vecUp, vecDown = xy{-1, 0}, xy{1, 0}, xy{0, -1}, xy{0, 1}
var allVecs = []xy{vecLeft, vecRight, vecUp, vecDown}

func movePacman(sc *screen, vec xy) {
	if b, loc := sc.tryMove(sc.p.loc.add(vec)); b {
		sc.p.move(loc, vec)
	}
}

func moveMonster(sc *screen, m *monster) {
	for {
		vec := m.vec
		if m.time == 0 {
			vec = m.getNextVec(m, sc)
		}
		m.liveLife()

		if b, loc := sc.tryMove(m.loc.add(vec)); b {
			m.move(loc, vec)
			return
		}
	}
}

func getPacmanVec(vec xy) (xy, bool) {
	switch getChFromView() {
	case 'h':
		vec = vecLeft
	case 'j':
		vec = vecDown
	case 'k':
		vec = vecUp
	case 'l':
		vec = vecRight

	case 'q':
		return vec, true
	}
	return vec, false
}

func runGame(sc *screen) {
Game:
	for {
		vec, quitGame := getPacmanVec(sc.p.vec)
		if quitGame {
			break Game
		}
		movePacman(sc, vec)

		for i := 0; i < len(sc.ms); i++ {
			m := &sc.ms[i]
			moveMonster(sc, m)
		}

		if sc.pacmanMetMonster() {
			if sc.p.hasExtraPower() {
				// eat monster
				// reset monster location
			} else {
				updateView(sc)
				break Game // pacman died
			}
		}

		switch sc.stuffUnderPacman() {
		case food:
			sc.eatFood()
		case powerFood:
			sc.eatFood()
			sc.p.addExtraPower(100)
		}

		updateView(sc)
		if sc.dots == 0 {
			break Game // pacman ate all food
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func buildMonsters() []monster {
	random := monster{getNextVec: getNextVecRandomly, stablePeriod: 5}
	shortestPath := monster{getNextVec: getNextVecOfShortestPath, stablePeriod: 3}
	return []monster{random, shortestPath}
}

func main() {
	initView()
	defer disposeView()

	p := buildPacman()
	ms := buildMonsters()
	sc := buildScreen(p, ms)
	confViewCapable(&sc)

	runGame(&sc)
}
