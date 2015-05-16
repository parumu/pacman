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
	p          xy
	ms         []xy
	mat        [][]rune
	size       xy
	dots       int
	extraPower int
}

func canMoveTo(sc *screen, xy xy) bool {
	return sc.mat[xy.y][xy.x] != '+'
}

func (sc *screen) getMatCell(xy xy) rune {
	return sc.mat[xy.y][xy.x]
}

func (sc *screen) setMatCell(xy xy, c rune) {
	sc.mat[xy.y][xy.x] = c
}

func (sc *screen) stuffUnderPacman() string {
	switch sc.getMatCell(sc.p) {
	case '.':
		return "Food"
	case 'O':
		return "PowerFood"
	default:
		return "None"
	}
}

func (sc *screen) eatFood() {
	switch c := sc.getMatCell(sc.p); c {
	case '.', 'O':
		sc.dots--
		sc.setMatCell(sc.p, ' ')
	}
}

func (sc *screen) tryMovePacman(np xy) {
	if canMoveTo(sc, np) {
		sc.p = np
	}
	if sc.extraPower > 0 {
		sc.extraPower--
	}
}

func (sc *screen) hasExtraPower() bool {
	return sc.extraPower > 0
}

func (sc *screen) addExtraPower(i int) {
	sc.extraPower += i
}

func (sc *screen) hitMonster() bool {
	for _, m := range sc.ms {
		if m == sc.p {
			return true
		}
	}
	return false
}

func buildScreen() screen {
	s := `
++++++++++++++|
+            +|
+ ++++ +++++ +|
+O++++ +++++ +|
+ ++++ +++++ +|
+             |
+ ++++ + +++++|
+      +     +|
++++++ +++++ +|
++++++ +      |
++++++ + ++++*|
++++++ + ++++*|
******   ++++M|
++++++ + +++++|
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
	p, o := xy{}, xy{}
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
				p.y = row
				p.x = col
				*c = ' '
			case 'M':
				o.y = row
				o.x = col
				*c = ' '
			}
		}
	}

	return screen{
		p:          p,
		ms:         []xy{o, o, o, o},
		mat:        m,
		size:       xy{maxX, maxY},
		dots:       dots,
		extraPower: 0}
}
