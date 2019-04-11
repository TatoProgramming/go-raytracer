package main

import(
	"math/rand"
)

type Material interface {
	scatter(r *Ray, rec HitRecord) (bool, Vector, Ray)
}

type Metal struct {
  Albedo Vector
}

type Lambertian struct {
  Albedo Vector
}

type Dielectric struct {
  refractionIndex float64
}

func (l Lambertian) scatter(r *Ray, rec HitRecord) (bool, Vector, Ray) {
  target := rec.P.Add(rec.Normal).Add(randomInUnitSphere())
  scattered := Ray{rec.P, target.Subtract(rec.P)}
  return true, l.Albedo, scattered
}

func (l Metal) scatter(r *Ray, rec HitRecord) (bool, Vector, Ray) {
  unitVector := r.Direction.MakeUnitVector()
  reflected := unitVector.Reflect(rec.Normal)
  scattered := Ray{rec.P, reflected}
  return scattered.Direction.Dot(rec.Normal) > 0 , l.Albedo, scattered
}

func (l Dielectric) scatter(r *Ray, rec HitRecord) (bool, Vector, Ray) {
	var scattered Ray
	var outwardNormal Vector
	reflected := r.Direction.Reflect(rec.Normal)
	var ni_over_nt float64
	attenuation := Vector{1.0, 1.0, 1.0}

	var reflectProb float64
	var cosine float64
	var refracted Vector

	if r.Direction.Dot(rec.Normal) > 0{
		outwardNormal = Vector{-rec.Normal.X,-rec.Normal.Y,-rec.Normal.Z}
		ni_over_nt = l.refractionIndex
		cosine = l.refractionIndex * r.Direction.Dot(rec.Normal) / r.Direction.Length()
	}else{
		outwardNormal = rec.Normal
		ni_over_nt = 1.0 / l.refractionIndex
		cosine = -r.Direction.Dot(rec.Normal) / r.Direction.Length()
	}


	isRefracted, refracted := r.Direction.Refract(outwardNormal, ni_over_nt)
	if isRefracted {
		reflectProb = schlick(cosine, l.refractionIndex)
	}else{
		reflectProb = 1.0
	}
	if rand.Float64() < reflectProb {
		scattered = Ray{rec.P, reflected}
	}else{
		scattered = Ray{rec.P, refracted}
	}
	return true, attenuation, scattered
}
