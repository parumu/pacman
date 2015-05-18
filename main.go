package main

import (
	"time"

	"github.com/parumu/gocurses"
)

func movePacman(sc *screen, vec vec) {
	sc.p.liveLife()

	if b, loc := sc.tryMove(sc.p.loc.addVec(vec)); b {
		sc.p.move(loc, vec)
	}
}

func moveMonster(sc *screen, m *monster) {
	for i := 0; i < 10; i++ { // exit after 10 tries to avoid infinite loop
		vec := m.vec
		if m.time == 0 {
			if sc.p.hasExtraPower() {
				vec = m.getScatterVec(m, sc)
			} else {
				vec = m.getChaseVec(m, sc)
			}
		}
		m.liveLife()

		if b, loc := sc.tryMove(m.loc.addVec(vec)); b {
			m.move(loc, vec)
			return
		}
	}
}

func getPacmanVec(v vec) (vec, bool) {
	switch getChFromView() {
	case gocurses.KEY_LEFT:
		v = vecLeft
	case gocurses.KEY_DOWN:
		v = vecDown
	case gocurses.KEY_UP:
		v = vecUp
	case gocurses.KEY_RIGHT:
		v = vecRight

	case 'q':
		return v, true
	}
	return v, false
}

func runGame(sc *screen) {
	pacmanDied := func() bool {
		if b, m := sc.pacmanMetMonster(); b {
			if !sc.p.hasExtraPower() {
				return true
			}
			m.loc = sc.monstHome // ate monster
		}
		return false
	}

Game:
	for {
		vec, quitGame := getPacmanVec(sc.p.vec)
		if quitGame {
			break Game
		}
		movePacman(sc, vec)
		if pacmanDied() {
			break Game
		}

		for i := 0; i < len(sc.ms); i++ {
			m := &sc.ms[i]
			moveMonster(sc, m)
			if pacmanDied() {
				break Game
			}
		}

		switch sc.stuffUnderPacman() {
		case food:
			sc.eatFood()
		case powerFood:
			sc.eatFood()
			sc.p.addExtraPower(100)
			sc.resetAllMonstTimes()
		}

		updateView(sc)
		if sc.dots == 0 {
			break Game // pacman ate all food
		}
		time.Sleep(100 * time.Millisecond)
	}
	updateView(sc)
}

func buildMonsters() []monster {
	rand := func(i int) monster {
		return monster{
			name:          "Random",
			getChaseVec:   getNextVecRandomly,
			getScatterVec: getNextVecRandomly,
			stablePeriod:  i}
	}
	sp := monster{
		name:          "Shortest Path",
		getChaseVec:   getNextVecOfShortestPath,
		getScatterVec: getNextVecOfLongestPath,
		stablePeriod:  10}

	hv := monster{
		name:          "Horizontal-Vertical",
		getChaseVec:   getNextVecByHorVer,
		getScatterVec: getNextVecByHorVerRev,
		stablePeriod:  3}

	return []monster{
		rand(5),
		rand(3),
		sp,
		hv}
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
