package main

import "math"

type Vector struct {
	X, Y, Z float64
}

func NewVector(f float64) Vector {
	return Vector{f, f, f}
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
