package mango

import (
	"math"

	"github.com/EliCDavis/vector"
)

func BuildRing(
	sides int,
	height float64,
	bottomRadius float64,
	topRadius float64,
	bottomOffset vector.Vector3,
	topOffset vector.Vector3,
) Mesh {
	tris := make([]Tri, 0)
	vertices := make([]vector.Vector3, 0)

	angleIncrement := (2.0 * math.Pi) / float64(sides)
	for i := 0; i < sides; i++ {
		aergea := angleIncrement * float64(i)
		v := vector.NewVector3(math.Cos(aergea)*bottomRadius, 0, math.Sin(aergea)*bottomRadius).Add(bottomOffset)
		vTop := vector.NewVector3(math.Cos(aergea)*topRadius, 0, math.Sin(aergea)*topRadius).Add(bottomOffset).Add(topOffset)
		vertices = append(vertices, v, vTop.Add(vector.Vector3Up().MultByConstant(height)))
	}

	for i := 1; i < sides; i++ {
		curVert := i * 2
		tris = append(
			tris,
			NewTri(curVert+1, curVert, curVert-1),
			NewTri(curVert-1, curVert, curVert-2),
		)
	}
	tris = append(tris, NewTri(1, 0, (sides*2)-1))
	tris = append(tris, NewTri((sides*2)-1, 0, (sides*2)-2))

	return NewMesh(vertices, tris)
}
