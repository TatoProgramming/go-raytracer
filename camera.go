package main

type Camera struct {
	lowerLeft, horizontal, vertical, origin Vector
}

func (c Camera) getRay(u, v float64) Ray {
	return Ray{
		c.origin,
		c.lowerLeft.Add(c.horizontal.Mult(u)).Add(c.vertical.Mult(v)).Subtract(c.origin),
	}
}
