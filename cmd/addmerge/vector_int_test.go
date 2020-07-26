package main

import (
	"testing"

	"github.com/EliCDavis/vector"
	"github.com/stretchr/testify/assert"
)

func TestVectorIntRoundsProperly(t *testing.T) {
	tests := map[string]struct {
		input vector.Vector3
		want  vectorInt
		pow   int
	}{
		"0,1,2=>0,1,2":         {input: vector.NewVector3(0, 1, 2), want: vectorInt{0, 1, 2}, pow: 0},
		"0.1,1.1,2.1=>0,1,2":   {input: vector.NewVector3(0.1, 1.1, 2.1), want: vectorInt{0, 1, 2}, pow: 0},
		"0.1,1.1,2.1=>1,11,21": {input: vector.NewVector3(0.1, 1.1, 2.1), want: vectorInt{1, 11, 21}, pow: 1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, vector3ToInt(tc.input, tc.pow))
		})
	}
}
