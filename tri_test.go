package mango

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriValidity(t *testing.T) {
	tests := map[string]struct {
		input Tri
		want  bool
	}{
		"valid":            {input: NewTri(0, 1, 2), want: true},
		"first 2 match":    {input: NewTri(0, 0, 2), want: false},
		"last 2 match":     {input: NewTri(0, 1, 1), want: false},
		"first last match": {input: NewTri(2, 1, 2), want: false},
		"all same":         {input: NewTri(0, 0, 0), want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.input.Valid())
		})
	}
}
