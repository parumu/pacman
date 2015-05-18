package main

type xy struct {
	x, y int
}

func (x *xy) add(another xy) xy {
	return xy{x.x + another.x, x.y + another.y}
}

func (x *xy) addVec(vec vec) xy {
	return xy{x.x + vec.x, x.y + vec.y}
}
