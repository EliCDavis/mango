package mango

type Tri struct {
	p1 int
	p2 int
	p3 int
}

func NewTri(p1, p2, p3 int) Tri {
	return Tri{p1, p2, p3}
}
