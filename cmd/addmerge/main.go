package main

import (
	"os"

	"github.com/EliCDavis/mango"
	"github.com/EliCDavis/vector"
)

// MergeByRounding takes two meshes and welds them together to a certain decimal place.
func MergeByRounding(m1 mango.Mesh, m2 mango.Mesh, decimalPlace int) mango.Mesh {
	// =================== Finding unique vertices ============================
	// Vertices from both meshes
	allVerts := append(m1.Vertices(), m2.Vertices()...)

	// Mapping from rounded vector to index of original vertex in allVerts
	vertLU := make(map[vectorInt]int)

	// Mapping from rounded vector to whether or not it get's used by a triangle
	// in the resulting mesh
	vertLUUsed := make(map[vectorInt]bool)

	// count of unique vertices once rounded
	uniqueVertCount := 0
	filteredVertices := make([]vector.Vector3, 0)
	for _, v := range allVerts {
		vInt := vector3ToInt(v, decimalPlace)

		if _, ok := vertLU[vInt]; !ok {
			vertLU[vInt] = uniqueVertCount
			vertLUUsed[vInt] = false
			filteredVertices = append(filteredVertices, v)
			uniqueVertCount++
		}
	}

	// Building tris from unique vertices
	newTris := make([]mango.Tri, 0)
	for _, tri := range m1.Triangles() {
		v1 := vector3ToInt(m1.Vertices()[tri.P1()], decimalPlace)
		v2 := vector3ToInt(m1.Vertices()[tri.P2()], decimalPlace)
		v3 := vector3ToInt(m1.Vertices()[tri.P3()], decimalPlace)
		rebuilt := mango.NewTri(
			vertLU[v1],
			vertLU[v2],
			vertLU[v3],
		)
		if rebuilt.Valid() {
			vertLUUsed[v1] = true
			vertLUUsed[v2] = true
			vertLUUsed[v3] = true
			newTris = append(newTris, rebuilt)
		}
	}

	// Do the same again with tris from the other mesh
	for _, tri := range m2.Triangles() {
		v1 := vector3ToInt(m2.Vertices()[tri.P1()], decimalPlace)
		v2 := vector3ToInt(m2.Vertices()[tri.P2()], decimalPlace)
		v3 := vector3ToInt(m2.Vertices()[tri.P3()], decimalPlace)
		rebuilt := mango.NewTri(
			vertLU[v1],
			vertLU[v2],
			vertLU[v3],
		)
		if rebuilt.Valid() {
			vertLUUsed[v1] = true
			vertLUUsed[v2] = true
			vertLUUsed[v3] = true
			newTris = append(newTris, rebuilt)
		}
	}

	finalVerts := make([]vector.Vector3, 0)
	shiftBy := make([]int, len(filteredVertices))
	curShift := 0
	for vertIndex, v := range filteredVertices {
		vInt := vector3ToInt(v, decimalPlace)

		if vertLUUsed[vInt] {
			finalVerts = append(finalVerts, v)
		} else {
			// Not used, need to shift triangles who's points point to vertices that come after this unsed one
			curShift++
		}
		shiftBy[vertIndex] = curShift
	}

	// Shift all the triangles appropriately since we just removed a bunch of vertices no longer used
	finalTris := make([]mango.Tri, len(newTris))
	for triIndex, tri := range newTris {
		finalTris[triIndex] = mango.NewTri(
			tri.P1()-shiftBy[tri.P1()],
			tri.P2()-shiftBy[tri.P2()],
			tri.P3()-shiftBy[tri.P3()],
		)
	}

	return mango.NewMesh(finalVerts, finalTris)
}

func main() {

	// Un-comment to see vertices get merged/welded together
	// bottomLayer := mango.BuildRing(30, 1, vector.Vector3Up().MultByConstant(0))
	// middleLayer := mango.BuildRing(24, 0.95, vector.Vector3Up().MultByConstant(1))
	// topLayer := mango.BuildRing(30, 0.9, vector.Vector3Up().MultByConstant(2))
	// MergeByRounding(bottomLayer, MergeByRounding(middleLayer, topLayer, 0), 0).ToOBJ(os.Stdout)

	bottomLayer := mango.BuildRing(30, 3, vector.Vector3Up().MultByConstant(0))
	middleLayer := mango.BuildRing(24, 2, vector.Vector3Up().MultByConstant(1))
	topLayer := mango.BuildRing(16, 1, vector.Vector3Up().MultByConstant(2))
	bottomLayer.Add(middleLayer).Add(topLayer).ToOBJ(os.Stdout)
}
