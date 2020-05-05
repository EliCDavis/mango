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
		Tri{0, 1, 2},
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
