package main

type HitRecord struct {
	T         float64
	P, Normal Vector
}

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) PointAt(t float64) Vector {
	return r.Origin.Add(r.Direction.Mult(t))
}
