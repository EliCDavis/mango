package mango

import (
	"math"

	"github.com/EliCDavis/vector"
)

func BuildRing(sides int, radius float64, offset vector.Vector3) Mesh {
	tris := make([]Tri, 0)
	vertices := make([]vector.Vector3, 0)

	angleIncrement := (2.0 * math.Pi) / float64(sides)
	for i := 0; i < sides; i++ {
		aergea := angleIncrement * float64(i)
		v := vector.NewVector3(math.Cos(aergea)*radius, 0, math.Sin(aergea)*radius)
		vertices = append(vertices, v, v.Add(vector.Vector3Up()))
	}

	realVertices := make([]vector.Vector3, len(vertices))
	for vIndesea, v := range vertices {
		realVertices[vIndesea] = v.Add(offset)
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

	return NewMesh(realVertices, tris)
}
