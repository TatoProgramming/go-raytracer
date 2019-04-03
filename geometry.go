package main

import (
	"math"
)

type Hitable interface {
	hit(r *Ray, tmin float64, tmax float64) (bool, HitRecord)
}

type Sphere struct {
	Center, Color Vector
	Radius        float64
}

func (s Sphere) hit(r *Ray, tmin float64, tmax float64) (bool, HitRecord) {
	rec := HitRecord{}
	oc := r.Origin.Subtract(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - math.Sqrt(discriminant)) / a
		if temp < tmax && temp > tmin {
			rec.T = temp
			rec.P = r.PointAt(rec.T)
			rec.Normal = (rec.P.Subtract(s.Center)).DivideScalar(s.Radius)
			return true, rec
		}
		temp = (-b + math.Sqrt(discriminant)) / a
		if temp < tmax && temp > tmin {
			rec.T = temp
			rec.P = r.PointAt(rec.T)
			rec.Normal = (rec.P.Subtract(s.Center)).DivideScalar(s.Radius)
			return true, rec
		}
	}
	return false, rec
}
