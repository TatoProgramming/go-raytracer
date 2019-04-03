package main

import (
	"fmt"
	// "math"
	"os"
)

func color(r Ray) Vector {
	sphere := Sphere{Vector{Z: -1.0}, Vector{X: 1.0, Y: 0.0, Z: 0.0}, 0.5}
	if hitSphere(sphere, &r) {
		return sphere.Color
	}
	// t := hitSphere(sphere, &r)
	// if t > 0.0 {
	// n := r.PointAt(t).Subtract(Vector{Z: -1.0}).MakeUnitVector()
	// return Vector{n.X + 1, n.Y + 1, n.Z + 1}.Mult(.5)
	// }
	unitVector := r.Direction.MakeUnitVector()
	t := .5 * (unitVector.Y + 1.0)
	return NewVector(1.0).Mult(1.0 - t).Add(Vector{.5, .7, 1.}.Mult(t))
}

func hitSphere(s Sphere, r *Ray) bool {
	oc := r.Origin.Subtract(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
}

func main() {
	const colorFactor = 255.99
	nx, ny := 300, 150
	f, _ := os.Create("out.ppm")

	fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)

	lowerLeftCorner := Vector{-2.0, -1.0, -1.0}
	horizontal := Vector{X: 4.0}
	vertical := Vector{Y: 2.0}
	origin := NewVector(0.0)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := Ray{origin, lowerLeftCorner.Add(horizontal.Mult(u).Add(vertical.Mult(v)))}
			col := color(r)
			ir, ig, ib := int(col.X*colorFactor), int(col.Y*colorFactor), int(col.Z*colorFactor)
			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}

	f.Close()
	ConvertToPNG("out")
}
