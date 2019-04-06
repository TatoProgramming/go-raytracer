package main

import "math"

type Vector struct {
	X, Y, Z float64
}

func NewVector(f float64) Vector {
	return Vector{f, f, f}
}

func (a Vector) MultV(b Vector) Vector {
	return Vector{
		a.X * b.X,
		a.Y * b.Y,
		a.Z * b.Z,
	}
}

func (a Vector) Mult(t float64) Vector {
	return Vector{
		a.X * t,
		a.Y * t,
		a.Z * t,
	}
}

func (a Vector) Add(b Vector) Vector {
	return Vector{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

func (a Vector) Subtract(b Vector) Vector {
	return Vector{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
	}
}

func (a Vector) Divide(b Vector) Vector {
	return Vector{
		a.X / b.X,
		a.Y / b.Y,
		a.Z / b.Z,
	}
}

func (a Vector) DivideScalar(t float64) Vector {
	return Vector{
		a.X / t,
		a.Y / t,
		a.Z / t,
	}
}

func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.SquaredLength())
}

func (v Vector) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector) MakeUnitVector() Vector {
	return v.DivideScalar(v.Length())
}

func (v Vector) Reflect(n Vector) Vector {
	return v.Subtract(n.Mult(v.Dot(n)).Mult(2.0))
}


func (v Vector) Refract(n Vector, ni_over_nt float64) (bool, Vector){
	uv := v.MakeUnitVector()
	dt := uv.Dot(n)
	discriminant := 1.0 - ni_over_nt*ni_over_nt*(1.0 - dt*dt)
	var refracted Vector
	if discriminant > 0.0{
		refracted = uv.Subtract(n.Mult(dt)).Mult(ni_over_nt).Subtract(n.Mult(math.Sqrt(discriminant)))
		return true, refracted
	}
	return false, refracted
}
