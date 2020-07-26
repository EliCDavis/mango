package mango

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/EliCDavis/vector"
	"github.com/stretchr/testify/assert"
)

func TestObj(t *testing.T) {
	// ARRANGE ======================================================
	vertices := []vector.Vector3{
		vector.Vector3Zero(),
		vector.Vector3Right(),
		vector.Vector3Up(),
	}
	triangles := []Tri{
		{0, 1, 2},
	}
	m := Mesh{vertices, triangles}
	out := bytes.Buffer{}

	// ACT ==========================================================
	err := m.ToOBJ(&out)

	// ASSERT =======================================================
	assert.NoError(t, err)

	scanner := bufio.NewScanner(&out)

	scanner.Scan()
	assert.Equal(t, "v 0.00000 0.00000 0.00000", scanner.Text())
	scanner.Scan()
	assert.Equal(t, "v 1.00000 0.00000 0.00000", scanner.Text())
	scanner.Scan()
	assert.Equal(t, "v 0.00000 1.00000 0.00000", scanner.Text())
	scanner.Scan()
	assert.Equal(t, "f 1 2 3", scanner.Text())
}

func TestAdd(t *testing.T) {
	// ARRANGE ======================================================
	m1 := NewMesh([]vector.Vector3{
		vector.Vector3Up(),
		vector.Vector3Up().Add(vector.Vector3Right()),
		vector.Vector3Right(),
	}, []Tri{
		{0, 1, 2},
	})

	m2 := NewMesh([]vector.Vector3{
		vector.Vector3Up(),
		vector.Vector3Right(),
		vector.Vector3Zero(),
	}, []Tri{
		{0, 1, 2},
	})

	// ACT ==========================================================
	addedMesh := m1.Add(m2)

	// ASSERT =======================================================
	assert.Len(t, addedMesh.triangles, 2)
	assert.Equal(t, Tri{0, 1, 2}, addedMesh.triangles[0])
	assert.Equal(t, Tri{3, 4, 5}, addedMesh.triangles[1])

	assert.Len(t, addedMesh.vertices, 6)
	assert.Equal(t, vector.Vector3Up(), addedMesh.vertices[0])
	assert.Equal(t, vector.Vector3Up().Add(vector.Vector3Right()), addedMesh.vertices[1])
	assert.Equal(t, vector.Vector3Right(), addedMesh.vertices[2])
	assert.Equal(t, vector.Vector3Up(), addedMesh.vertices[3])
	assert.Equal(t, vector.Vector3Right(), addedMesh.vertices[4])
	assert.Equal(t, vector.Vector3Zero(), addedMesh.vertices[5])
}
