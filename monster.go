package main

import (
	"math"
	"math/rand"
)

type monster struct {
	name          string
	loc           xy
	vec           vec
	getChaseVec   func(m *monster, sc *screen) vec
	getScatterVec func(m *monster, sc *screen) vec
	time          int
	stablePeriod  int
}

func (m *monster) move(loc xy, vec vec) {
	m.loc = loc
	m.vec = vec
}

func (m *monster) liveLife() {
	m.time = (m.time + 1) % m.stablePeriod
}

func getNextVecByHorVer(m *monster, sc *screen) vec {
	if rand.Intn(2)%2 == 0 {
		if sc.p.loc.x < m.loc.x {
			return vecLeft
		}
		return vecRight
	}
	if sc.p.loc.y < m.loc.y {
		return vecUp
	}
	return vecDown
}

func getNextVecByHorVerRev(m *monster, sc *screen) vec {
	vec := getNextVecByHorVer(m, sc)
	return oppositeVec(vec)
}

func getNextVecRandomly(_ *monster, _ *screen) vec {
	return allVecs[rand.Intn(4)]
}

func getShortestDist2Pacman(loc xy, sc *screen) (bool, int) {
	matBak := sc.cloneMat()
	defer (func() { sc.mat = matBak })()

	q := new(queue)
	q.push(loc)
	dist := 0

	for !q.isEmpty() {
		dist++

		loc := q.pop().(xy)
		if sc.p.loc == loc {
			return true, dist
		}
		sc.mat[loc.y][loc.x] = '+'

		for _, vec := range allVecs {
			if b, nextLoc := sc.tryMove(loc.addVec(vec)); b {
				q.push(nextLoc)
			}
		}
	}
	return false, math.MaxInt16
}

func getNextVecOfXestPath(m *monster, sc *screen, pred func(int, int) bool, distInit int) vec {
	retVec, retDist := vec{}, distInit

	for _, vec := range allVecs {
		canMove, loc := sc.tryMove(m.loc.addVec(vec))
		if !canMove {
			continue
		}
		if isReachable, dist := getShortestDist2Pacman(loc, sc); isReachable {
			if pred(dist, retDist) {
				retDist = dist
				retVec = vec
			}
		}
	}
	return retVec
}

func getNextVecOfShortestPath(m *monster, sc *screen) vec {
	pred := func(a int, b int) bool { return a < b }
	return getNextVecOfXestPath(m, sc, pred, math.MaxInt16)
}

func getNextVecOfLongestPath(m *monster, sc *screen) vec {
	pred := func(a int, b int) bool { return a > b }
	return getNextVecOfXestPath(m, sc, pred, -1)
}
