package main

import (
	"math"

	"github.com/EliCDavis/vector"
)

type vectorInt struct {
	x int
	y int
	z int
}

func vector3ToInt(v vector.Vector3, power int) vectorInt {
	newPower := math.Pow10(power)
	return vectorInt{
		x: int(math.Round(v.X() * newPower)),
		y: int(math.Round(v.Y() * newPower)),
		z: int(math.Round(v.Z() * newPower)),
	}
}

func (v vectorInt) ToRegularVector() vector.Vector3 {
	return vector.NewVector3(
		float64(v.x),
		float64(v.y),
		float64(v.z),
	)
}
