package mango

import (
	"fmt"
	"io"

	"github.com/EliCDavis/vector"
)

// Mesh is a collection of geometry data that can be rendered
type Mesh struct {
	vertices  []vector.Vector3
	triangles []Tri
}

// NewMesh creates a new mesh
func NewMesh(vertices []vector.Vector3, triangles []Tri) Mesh {
	return Mesh{vertices, triangles}
}

var emptyMesh = Mesh{}

// NewEmptyMesh creates a new mesh with no faces or triangles
func NewEmptyMesh(vertices []vector.Vector3, triangles []Tri) Mesh {
	return emptyMesh
}

// Vertices returns all vertices to the mesh
func (m Mesh) Vertices() []vector.Vector3 {
	return m.vertices
}

// Triangles returns all triangles to the mesh
func (m Mesh) Triangles() []Tri {
	return m.triangles
}

// Add takes two meshes and appends their data together in an unwelded fashion.
func (m Mesh) Add(m2 Mesh) Mesh {
	newVertices := append(m.vertices, m2.vertices...)
	newTris := m.triangles

	m1VertCount := len(m.vertices)
	for _, tri := range m2.triangles {
		newTris = append(newTris, NewTri(
			tri.p1+m1VertCount,
			tri.p2+m1VertCount,
			tri.p3+m1VertCount,
		))
	}

	return Mesh{newVertices, newTris}
}

// ToOBJ writes out the mesh in obj format
func (m Mesh) ToOBJ(w io.Writer) error {
	for _, v := range m.vertices {
		io.WriteString(w, fmt.Sprintf("v %.5f %.5f %.5f\n", v.X(), v.Y(), v.Z()))
	}

	for _, t := range m.triangles {
		io.WriteString(w, fmt.Sprintf("f %d %d %d\n", t.p1+1, t.p2+1, t.p3+1))
	}

	return nil
}
