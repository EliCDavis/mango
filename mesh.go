package mango

import (
	"fmt"
	"io"

	"github.com/EliCDavis/vector"
)

type Mesh struct {
	vertices  []vector.Vector3
	triangles []Tri
}

func NewMesh(vertices []vector.Vector3, triangles []Tri) Mesh {
	return Mesh{vertices, triangles}
}

func (m Mesh) Add(m2 Mesh) Mesh {
	resultingTriangles := m.triangles
	for _, t := range m2.triangles {
		resultingTriangles = append(resultingTriangles, Tri{
			p1: t.p1 + len(m.triangles),
			p2: t.p2 + len(m.triangles),
			p3: t.p3 + len(m.triangles),
		})
	}

	return Mesh{
		vertices:  append(m.vertices, m2.vertices...),
		triangles: resultingTriangles,
	}
}

func (m Mesh) ToOBJ(w io.Writer) error {
	for _, v := range m.vertices {
		io.WriteString(w, fmt.Sprintf("v %.5f %.5f %.5f\n", v.X(), v.Y(), v.Z()))
	}

	for _, t := range m.triangles {
		io.WriteString(w, fmt.Sprintf("f %d %d %d\n", t.p1+1, t.p2+1, t.p3+1))
	}

	return nil
}
