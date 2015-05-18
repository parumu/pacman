package main

type vec xy

var vecLeft, vecRight, vecUp, vecDown = vec{-1, 0}, vec{1, 0}, vec{0, -1}, vec{0, 1}
var allVecs = [4]vec{vecLeft, vecRight, vecUp, vecDown}

func oppositeVec(v vec) vec {
	switch v {
	case vecLeft:
		return vecRight
	case vecRight:
		return vecLeft
	case vecUp:
		return vecDown
	default:
		return vecUp
	}
}

func (v vec) String() string {
	switch v {
	case vecLeft:
		return "Left"
	case vecRight:
		return "Right"
	case vecUp:
		return "Up"
	default:
		return "Down"
	}
}
