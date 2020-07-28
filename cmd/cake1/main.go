package main

import (
	"math"
	"os"

	"github.com/EliCDavis/mango"
	"github.com/EliCDavis/vector"
)

func Cylinder(
	sides int,
	height float64,
	bottomRadius float64,
	topRadius float64,
	bottomOffset vector.Vector3,
	topOffset vector.Vector3,
) mango.Mesh {
	shaft := mango.BuildRing(sides, height, bottomRadius, topRadius, bottomOffset, topOffset)
	top := mango.BuildRing(
		sides,
		0,
		topRadius,
		0,
		bottomOffset.
			Add(topOffset).
			Add(vector.NewVector3(0, height, 0)),
		vector.Vector3Zero(),
	)
	bottom := mango.BuildRing(
		sides,
		0,
		0,
		bottomRadius,
		bottomOffset,
		vector.Vector3Zero(),
	)
	return shaft.Add(top).Add(bottom)
}

func ThiccRing(
	sides int,
	height float64,
	bottomRadius float64,
	topRadius float64,
	bottomOffset vector.Vector3,
	topOffset vector.Vector3,
	thickness float64,
) mango.Mesh {

	shaft := mango.BuildRing(sides, height, bottomRadius, topRadius, bottomOffset, topOffset)

	top := mango.BuildRing(
		sides,
		0,
		topRadius,
		topRadius-thickness,
		bottomOffset.
			Add(topOffset).
			Add(vector.NewVector3(0, height, 0)),
		vector.Vector3Zero(),
	)

	shaftDown := mango.BuildRing(
		sides,
		-height,
		bottomRadius-thickness,
		topRadius-thickness,
		bottomOffset.Add(vector.Vector3Up().MultByConstant(height)),
		topOffset,
	)

	return shaft.Add(top).Add(shaftDown)
}

func CakeTier(radiues, height float64, offset vector.Vector3) mango.Mesh {
	thiccConst := .15

	cakeSides := 30

	plate := Cylinder(
		cakeSides,
		thiccConst,
		radiues,
		radiues,
		vector.Vector3Up().MultByConstant(0).Add(offset),
		vector.Vector3Zero(),
	)

	cake := Cylinder(
		cakeSides,
		height-thiccConst,
		radiues-thiccConst,
		radiues-thiccConst,
		vector.Vector3Up().MultByConstant(thiccConst).Add(offset),
		vector.Vector3Zero(),
	)

	icing := ThiccRing(
		cakeSides,
		thiccConst,
		radiues-thiccConst-thiccConst,
		radiues-thiccConst-thiccConst,
		vector.Vector3Up().MultByConstant(height).Add(offset),
		vector.Vector3Zero(),
		thiccConst,
	)

	return plate.Add(cake).Add(icing)
}

func CakeColumn(radiues, height float64, numColumns int, offset vector.Vector3) mango.Mesh {
	resultingColumns := mango.NewEmptyMesh()

	columnSideConstant := 18

	angle := (math.Pi * 2.0) / float64(numColumns)
	for i := 0; i < numColumns; i++ {
		resultingColumns = resultingColumns.Add(Cylinder(
			columnSideConstant,
			height,
			.1,
			.1,
			offset.Add(vector.NewVector3(
				math.Sin(angle*float64(i))*radiues,
				0,
				math.Cos(angle*float64(i))*radiues,
			)),
			vector.Vector3Zero(),
		))
	}

	return resultingColumns
}

func main() {
	bottomLayer := CakeTier(3.0, 1.0, vector.Vector3Zero())
	bottomColumns := CakeColumn(1.7, .5, 6, vector.Vector3Up())
	middleLayer := CakeTier(2.1, 1.0, vector.Vector3Up().MultByConstant(1.5))
	middleColumns := CakeColumn(1, .5, 4, vector.Vector3Up().MultByConstant(2.5))
	topLayer := CakeTier(1.4, 1.0, vector.Vector3Up().MultByConstant(3))

	bottomLayer.
		Add(bottomColumns).
		Add(middleLayer).
		Add(middleColumns).
		Add(topLayer).
		ToOBJ(os.Stdout)
}
