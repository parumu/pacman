package main

import (
	"time"

	"github.com/parumu/gocurses"
)

var vecLeft, vecRight, vecUp, vecDown = xy{-1, 0}, xy{1, 0}, xy{0, -1}, xy{0, 1}

func movePacman(sc *screen, vec xy) {
	np := sc.p.add(vec)

	// handle entering warp hole
	if np.x == -1 {
		np.x = sc.size.x - 1
	} else if np.x == sc.size.x {
		np.x = 0
	}

	sc.tryMovePacman(np)
}

func moveMonsters(sc *screen) {

}

func getUserInput(vec xy) (xy, bool) {
	switch gocurses.Stdscr.Getch() {
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
	vec, quitGame := vecLeft, false
Game:
	for {
		vec, quitGame = getUserInput(vec)
		if quitGame {
			break Game
		}
		movePacman(sc, vec)
		moveMonsters(sc)

		if sc.hitMonster() {
			if sc.hasExtraPower() {
				// eat monster
				// reset monster location
			} else {
				break Game // pacman died
			}
		}

		switch sc.stuffUnderPacman() {
		case food:
			sc.eatFood()
		case powerFood:
			sc.eatFood()
			sc.addExtraPower(100)
		}

		updateView(sc)
		if sc.dots == 0 {
			break Game // pacman ate all food
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	initView()
	defer disposeView()

	sc := buildScreen()
	confViewCapable(&sc)

	runGame(&sc)
}
