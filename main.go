package main

import (
	"fmt"
	"math"
	"os"
)

func color(r *Ray, world *Scene) Vector {
	isHit, hitRec := world.hit(r, 0.0, math.MaxFloat64)
	if isHit {
		return Vector{hitRec.Normal.X + 1, hitRec.Normal.Y + 1, hitRec.Normal.Z + 1}.Mult(.5)
	}
	unitVector := r.Direction.MakeUnitVector()
	t := .5 * (unitVector.Y + 1.0)
	return NewVector(1.0).Mult(1.0 - t).Add(Vector{.5, .7, 1.}.Mult(t))
}

func main() {
	const colorFactor = 255.99
	nx, ny, ns := 300, 150, 100
	f, _ := os.Create("out.ppm")

	fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)

	lowerLeftCorner := Vector{-2.0, -1.0, -1.0}
	horizontal := Vector{X: 4.0}
	vertical := Vector{Y: 2.0}
	origin := NewVector(0.0)

	world := Scene{make([]Hitable, 2)}
	world.Objects[0] = Sphere{Center: Vector{Z: -1.0}, Color: Vector{X: 1.0, Y: 0.0, Z: 0.0}, Radius: 0.5}
	world.Objects[1] = Sphere{Center: Vector{Y: -100.5, Z: -1.0}, Color: Vector{X: 1.0, Y: 0.0, Z: 0.0}, Radius: 100.0}

	cam := Camera{
		Vector{-2.0, -1.0, -1.0},
		Vector{X: 4.0},
		Vector{Y: 2.0},
		NewVector(0.0),
	}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := Vector
			for s := 0; s < ns;
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := Ray{origin, lowerLeftCorner.Add(horizontal.Mult(u).Add(vertical.Mult(v)))}
			col.Add(color(&r, &world))
			ir, ig, ib := int(col.X*colorFactor), int(col.Y*colorFactor), int(col.Z*colorFactor)
			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}

	f.Close()
	ConvertToPNG("out")
}
