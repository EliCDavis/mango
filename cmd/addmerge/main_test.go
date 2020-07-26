package main

import (
	"testing"

	"github.com/EliCDavis/vector"
	"github.com/stretchr/testify/assert"
)

func TestMergeByRounding(t *testing.T) {
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
	addedMesh := m1.MergeByRounding(m2, 4)

	// ASSERT =======================================================
	assert.Len(t, addedMesh.triangles, 2)
	assert.Equal(t, Tri{0, 1, 2}, addedMesh.triangles[0])
	assert.Equal(t, Tri{0, 2, 3}, addedMesh.triangles[1])

	assert.Len(t, addedMesh.vertices, 4)
	assert.Equal(t, vector.Vector3Up(), addedMesh.vertices[0])
	assert.Equal(t, vector.Vector3Up().Add(vector.Vector3Right()), addedMesh.vertices[1])
	assert.Equal(t, vector.Vector3Right(), addedMesh.vertices[2])
	assert.Equal(t, vector.Vector3Zero(), addedMesh.vertices[3])
}

func TestMergeByRoundingRemoveUnusedVertices(t *testing.T) {
	// ARRANGE ======================================================
	// Really weird triangle that gets collapsed and now theres unused vertices
	m1 := NewMesh([]vector.Vector3{
		vector.NewVector3(0, 1, 0),
		vector.NewVector3(0, 1.01, 0),
		vector.NewVector3(6, 1.01, 0),
	}, []Tri{
		{0, 1, 2},
	})

	m2 := NewMesh([]vector.Vector3{
		vector.Vector3Up(),
		vector.Vector3Up().Add(vector.Vector3Right()),
		vector.Vector3Right(),
	}, []Tri{
		{0, 1, 2},
	})

	// ACT ==========================================================
	addedMesh := m1.MergeByRounding(m2, 1)

	// ASSERT =======================================================
	assert.Len(t, addedMesh.triangles, 1)
	assert.Equal(t, Tri{0, 1, 2}, addedMesh.triangles[0])

	assert.Len(t, addedMesh.vertices, 3)
	assert.Equal(t, vector.Vector3Up(), addedMesh.vertices[0])
	assert.Equal(t, vector.Vector3Up().Add(vector.Vector3Right()), addedMesh.vertices[1])
	assert.Equal(t, vector.Vector3Right(), addedMesh.vertices[2])
}
