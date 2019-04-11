package main

import(
	"math/rand"
)

type Scene struct {
	Objects []Hitable
	Size int
}

func (scene Scene) hit(r *Ray, tmin float64, tmax float64) (bool, HitRecord) {
	var finalRec, tempRecord HitRecord
	var isHit bool
	hitAnything := false
	closest := tmax
	for i := 0; i < scene.Size; i++ {
		isHit, tempRecord = scene.Objects[i].hit(r, tmin, closest)
		if isHit {
			hitAnything = true
			closest = tempRecord.T
			finalRec = tempRecord
		}
	}
	return hitAnything, finalRec
}

func RandomScene() Scene {
	const n = 500
	list := make([]Hitable, n + 1)

	list[0] = Sphere{
		Vector{0.0, -1000.0, 0.0},
		1000.0,
		Lambertian{Vector{0.5, 0.5, 0.5}},
	}
	i := 1
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat :=  rand.Float64()
			center := Vector{float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64()}
			if center.Subtract(Vector{4.0, 0.2, 0.0}).Length() > 0.9 {
				if chooseMat < 0.8 { // diffuse
					list[i] = Sphere{
						center,
						0.2,
						Lambertian{Vector{rand.Float64()*rand.Float64(), rand.Float64()*rand.Float64(), rand.Float64()*rand.Float64()}},
					}
					i++
				} else if chooseMat < 0.95 {
					list[i] = Sphere{
						center,
						0.2,
						Metal{
							Vector{0.5*(1.0+ rand.Float64()), 0.5*(1.0+ rand.Float64()), 0.5*(1.0+ rand.Float64())},
						},
					}
					i++
				}else {
					list[i] = Sphere{
						center,
						0.2,
						Dielectric{1.5},
					}
					i++
				}
			}
		}
	}

	list[i] = Sphere{
		Vector{0.0, 1.0, 0.0},
		1.0,
		Dielectric{1.5},
	}
	i++
	list[i] = Sphere{
		Vector{-4.0, 1.0, 0.0},
		1.0,
		Lambertian{ Vector{0.4, 0.2, 0.1} },
	}
	i++
	list[i] = Sphere{
		Vector{4.0, 1.0, 0.0},
		1.0,
		Lambertian{ Vector{0.7, 0.6, 0.5} },
	}
	i++

	return Scene{list, i}
}
