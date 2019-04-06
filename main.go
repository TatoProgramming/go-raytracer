package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

func color(r *Ray, world *Scene, depth int) Vector {
	isHit, hitRec := world.hit(r, 0.001, math.MaxFloat64)
	if isHit {
		isReflected, attenuation, scattered := hitRec.Mat.scatter(r, hitRec)
		if depth < 50 && isReflected {
			return attenuation.MultV(color(&scattered, world, depth+1))
		}else{
			return NewVector(0.0)
		}
		// target := hitRec.P.Add(hitRec.Normal).Add(randomInUnitSphere())
		// randRay := Ray{hitRec.P, target.Subtract(hitRec.P)}
		// return color(&randRay, world).Mult(.5)
	}
	unitVector := r.Direction.MakeUnitVector()
	t := .5 * (unitVector.Y + 1.0)
	return NewVector(1.0).Mult(1.0 - t).Add(Vector{.5, .7, 1.}.Mult(t))
}

func randomInUnitSphere() Vector {
	var p Vector
	for valid := true; valid; valid = p.SquaredLength() >= 1.0 {
		p = Vector{rand.Float64(), rand.Float64(), rand.Float64()}.Mult(2.0).Subtract(NewVector(1.0))
	}
	return p
}

func main() {
	const colorFactor = 255.99
	nx, ny, ns := 720/2, 480/2, 100
	f, _ := os.Create("out.ppm")

	fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)

	world := Scene{make([]Hitable, 4)}
	world.Objects[0] = Sphere{
		Center: Vector{Z: -1.0},
		Color: Vector{X: 1.0, Y: 0.0, Z: 0.0},
		Radius: 0.5,
		Mat: Lambertian{
			Albedo: Vector{0.8,0.3,0.3},
		},
	}

	world.Objects[1] = Sphere{
		Center: Vector{Y: -100.5, Z: -1.0},
		Color: Vector{X: 1.0, Y: 0.0, Z: 0.0},
		Radius: 100.0,
		Mat: Lambertian{
			Albedo: Vector{0.8,0.8,0.0},
		},
	}

	world.Objects[2] = Sphere{
		Center: Vector{1.0, 0.0, -1.0},
		Color: Vector{X: 1.0, Y: 0.0, Z: 0.0},
		Radius: 0.5,
		Mat: Metal{
			Albedo: Vector{0.8,0.6,0.2},
		},
	}

	world.Objects[3] = Sphere{
		Center: Vector{-1.0, 0.0, -1.0},
		Color: Vector{X: 1.0, Y: 0.0, Z: 0.0},
		Radius: 0.5,
		Mat: Dielectric{
			refractionIndex: 1.5,
		},
	}

	cam := Camera{
		Vector{-2.0, -1.0, -1.0},
		Vector{X: 4.0},
		Vector{Y: 2.0},
		NewVector(0.0),
	}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			// super sampling
			col := NewVector(0.0)
			for s := 0; s < ns; s++{
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := cam.getRay(u, v)
				col = col.Add(color(&r, &world, 0))
			}
			col = col.DivideScalar(float64(ns))
			col = Vector{math.Sqrt(col.X),math.Sqrt(col.Y),math.Sqrt(col.Z)}
			ir, ig, ib := int(col.X*colorFactor), int(col.Y*colorFactor), int(col.Z*colorFactor)
			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}

	f.Close()
	ConvertToPNG("out")
}
