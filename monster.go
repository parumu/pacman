package main

import (
	"math"
	"math/rand"
)

type monster struct {
	loc          xy
	vec          xy
	getNextVec   func(m *monster, sc *screen) xy
	time         int
	stablePeriod int
}

func (m *monster) move(loc xy, vec xy) {
	m.loc = loc
	m.vec = vec
}

func (m *monster) liveLife() {
	m.time = (m.time + 1) % m.stablePeriod
}

func getNextVecRandomly(_ *monster, _ *screen) xy {
	return [4]xy{vecDown, vecLeft, vecRight, vecUp}[rand.Intn(4)]
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
			if b, nextLoc := sc.tryMove(loc.add(vec)); b {
				q.push(nextLoc)
			}
		}
	}
	return false, math.MaxInt16
}

func getNextVecOfShortestPath(m *monster, sc *screen) xy {
	minVec, minDist := xy{}, math.MaxInt16

	for _, vec := range allVecs {
		canMove, loc := sc.tryMove(m.loc.add(vec))
		if !canMove {
			continue
		}
		if isReachable, dist := getShortestDist2Pacman(loc, sc); isReachable {
			if dist < minDist {
				minDist = dist
				minVec = vec
			}
		}
	}
	return minVec
}
