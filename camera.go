package main

import(
	"math"
	 "math/rand"
 )

type Camera struct {
	lowerLeft, horizontal, vertical, origin, u, v, w Vector
	lensRadius float64
}

func randomInUnitDisc() Vector {
	var p Vector
	for valid := true; valid; valid = p.Dot(p) >= 1.0 {
		p = Vector{rand.Float64(), rand.Float64(), 0.0}.Mult(2.0).Subtract(Vector{1.0, 1.0, 0.0})
	}
	return p
}

func NewCamera(lookFrom, lookAt, vup Vector, vfov, aspect, aperture, focusDist float64) Camera{
	var cam Camera
	cam.lensRadius = aperture / 2.0
	theta := vfov*math.Pi/180.0
	halfHeight := math.Tan(theta/2.0)
	halfWidth := aspect * halfHeight
	cam.origin = lookFrom
	cam.w = lookFrom.Subtract(lookAt).MakeUnitVector()
	cam.u = vup.Cross(cam.w).MakeUnitVector()
	cam.v = cam.w.Cross(cam.u)
	cam.lowerLeft = cam.origin.Subtract(cam.u.Mult(halfWidth*focusDist)).Subtract(cam.v.Mult(halfHeight*focusDist)).Subtract(cam.w.Mult(focusDist))
	cam.horizontal = cam.u.Mult(2.0*halfWidth*focusDist)
	cam.vertical = cam.v.Mult(2.0*halfHeight*focusDist)
	return cam
}

func (c Camera) getRay(s, t float64) Ray {
	rd := randomInUnitDisc().Mult(c.lensRadius)
	offset := c.u.Mult(rd.X).Add(c.v.Mult(rd.Y))
	return Ray{
		c.origin.Add(offset),
		c.lowerLeft.Add(c.horizontal.Mult(s)).Add(c.vertical.Mult(t)).Subtract(c.origin).Subtract( offset),
	}
}
