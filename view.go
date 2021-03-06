package main

import (
	"fmt"

	"github.com/parumu/gocurses"
)

func initView() {
	gocurses.Initscr()
	gocurses.Cbreak()
	gocurses.Noecho()
	gocurses.Stdscr.Keypad(true)
	gocurses.CursSet(0)
	gocurses.Timeout(0)
}

func updateView(sc *screen) {
	for y := 0; y < sc.size.y; y++ {
		for x := 0; x < sc.size.x; x++ {
			c := sc.mat[y][x]
			switch c {
			case '+':
				gocurses.Attron(gocurses.A_REVERSE)
			default:
				gocurses.Attroff(gocurses.A_REVERSE)
			}
			gocurses.Mvaddch(y, x, c)
		}
	}
	gocurses.Attroff(gocurses.A_REVERSE)
	gocurses.Attron(gocurses.A_BOLD)

	p := 'P'
	if sc.p.hasExtraPower() {
		p = 'S'
	}
	gocurses.Mvaddch(sc.p.loc.y, sc.p.loc.x, p)

	for i, m := range sc.ms {
		mo := 'M'
		if sc.p.hasExtraPower() {
			mo = 'W'
		}
		gocurses.Mvaddch(m.loc.y, m.loc.x, mo)
		desc := fmt.Sprintf("%2d x %2d: %s(%d/%d) %5s",
			m.loc.y, m.loc.x, m.name, m.time, m.stablePeriod, m.vec)
		gocurses.Mvaddstr(sc.size.y+3+i, 0, desc)
	}
	gocurses.Mvaddstr(sc.size.y, 0, "Use arrow keys to move packman. 'q' to quit game.")
	gocurses.Mvaddstr(sc.size.y+2, 0, fmt.Sprintf("Dots Left: %3d", sc.dots))
	gocurses.Refresh()
}

func disposeView() {
	gocurses.End()
}

func getChFromView() int {
	return gocurses.Stdscr.Getch()
}

func confViewCapable(sc *screen) {
	y, x := gocurses.Getmaxyx()
	reqY, reqX := sc.size.y+len(sc.ms)+3, sc.size.x
	if reqY > y || reqX > x {
		panic(fmt.Sprintf("Minumum screen size is %dx%d", reqX, reqY))
	}
}
