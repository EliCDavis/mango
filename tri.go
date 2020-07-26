package mango

type Tri struct {
	p1 int
	p2 int
	p3 int
}

// NewTri creates a new tri
func NewTri(p1, p2, p3 int) Tri {
	return Tri{p1, p2, p3}
}

// P1 is the first point on our triangle, which is an index to the vertices array of a mesh
func (t Tri) P1() int {
	return t.p1
}

// P2 is the second point on our triangle, which is an index to the vertices array of a mesh
func (t Tri) P2() int {
	return t.p2
}

// P3 is the third point on our triangle, which is an index to the vertices array of a mesh
func (t Tri) P3() int {
	return t.p3
}

// Valid determines whether or not the contains 3 unique vertices.
func (t Tri) Valid() bool {
	if t.p1 == t.p2 {
		return false
	}
	if t.p1 == t.p3 {
		return false
	}
	if t.p2 == t.p3 {
		return false
	}
	return true
}
