package main

import (
	"fmt"
	"strings"
)

const food = "Food"
const powerFood = "PowerFood"
const none = "None"

type xy struct {
	x, y int
}

func (a *xy) add(b xy) xy {
	return xy{a.x + b.x, a.y + b.y}
}

type screen struct {
	p    pacman
	ms   []monster
	mat  [][]rune
	size xy
	dots int
}

func (sc *screen) cloneMat() [][]rune {
	newMat := make([][]rune, len(sc.mat))
	for row := 0; row < len(sc.mat); row++ {
		newRow := make([]rune, len(sc.mat[row]))
		newMat[row] = newRow
		for col := 0; col < len(sc.mat[row]); col++ {
			newRow[col] = sc.mat[row][col]
		}
	}
	return newMat
}

// returns if loc is on path adjusting loc in warp hole case
func (sc *screen) tryMove(loc xy) (bool, xy) {
	// handle warp hole case
	if loc.x == -1 {
		loc.x = sc.size.x - 1
	} else if loc.x == sc.size.x {
		loc.x = 0
	}

	if sc.mat[loc.y][loc.x] == '+' {
		return false, loc
	}
	return true, loc
}

func (sc *screen) getMatCell(xy xy) rune {
	return sc.mat[xy.y][xy.x]
}

func (sc *screen) setMatCell(xy xy, c rune) {
	sc.mat[xy.y][xy.x] = c
}

func (sc *screen) stuffUnderPacman() string {
	switch sc.getMatCell(sc.p.loc) {
	case '.':
		return food
	case 'O':
		return powerFood
	default:
		return none
	}
}

func (sc *screen) eatFood() {
	switch c := sc.getMatCell(sc.p.loc); c {
	case '.', 'O':
		sc.dots--
		sc.setMatCell(sc.p.loc, ' ')
	}
}

func (sc *screen) pacmanMetMonster() bool {
	for _, m := range sc.ms {
		if m.loc == sc.p.loc {
			return true
		}
	}
	return false
}

func buildScreen(p pacman, ms []monster) screen {
	s := `
++++++++++++++|
+            +|
+O++++ +++++ +|
+             |
+ ++++ + +++++|
+      +     +|
++++++ +++++ +|
++++++ +      |
++++++ + ++++*|
******   ++++M|
++++++ + +++++|
++++++ +      |
++++++ + +++++|
+            +|
+ ++++ +++++ +|
+O   +       P|
++++ + + +++++|
+      +     +|
+ ++++++++++ +|
+             |
++++++++++++++|
`
	strRev := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	// build maze base
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	maxY := len(lines)
	m := make([][]rune, maxY)

	for row := 0; row < maxY; row++ {
		l := strings.TrimRight(lines[row], "|")
		l = l + strRev(l[0:len(l)-1])
		fmt.Printf("Adding '%s'\n", l)
		m[row] = []rune(l)
	}
	maxX := len(m[0])

	// deploy items
	pLoc, mLoc := xy{}, xy{}
	dots := 0

	for row := 0; row < maxY; row++ {
		for col := 0; col < maxX; col++ {
			switch c := &m[row][col]; *c {
			case ' ':
				*c = '.'
				dots++
			case 'O':
				dots++
			case '*':
				*c = ' '
			case 'P':
				pLoc.y = row
				pLoc.x = col
				*c = ' '
			case 'M':
				mLoc.y = row
				mLoc.x = col
				*c = ' '
			}
		}
	}

	// set starting locations
	p.loc = pLoc
	for i := 0; i < len(ms); i++ {
		ms[i].loc = mLoc
	}

	return screen{
		p:    p,
		ms:   ms,
		mat:  m,
		size: xy{maxX, maxY},
		dots: dots}
}
