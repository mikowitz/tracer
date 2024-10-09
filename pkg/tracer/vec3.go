package tracer

import "math"

type Vec3 [3]float64

func (u Vec3) Neg() Vec3 {
	return Vec3{-u[0], -u[1], -u[2]}
}

func (u Vec3) Add(v Vec3) Vec3 {
	return Vec3{
		u[0] + v[0],
		u[1] + v[1],
		u[2] + v[2],
	}
}

func (u Vec3) Sub(v Vec3) Vec3 {
	return Vec3{
		u[0] - v[0],
		u[1] - v[1],
		u[2] - v[2],
	}
}

func (u Vec3) Mul(t float64) Vec3 {
	return Vec3{u[0] * t, u[1] * t, u[2] * t}
}

func (u Vec3) Prod(v Vec3) Vec3 {
	return Vec3{
		u[0] * v[0],
		u[1] * v[1],
		u[2] * v[2],
	}
}

func (u Vec3) Div(t float64) Vec3 {
	return u.Mul(1.0 / t)
}

func (u Vec3) Length() float64 {
	return math.Sqrt(u.LengthSquared())
}

func (u Vec3) LengthSquared() float64 {
	return u[0]*u[0] + u[1]*u[1] + u[2]*u[2]
}

func (u Vec3) Dot(v Vec3) float64 {
	return u[0]*v[0] + u[1]*v[1] + u[2]*v[2]
}

func (u Vec3) Cross(v Vec3) Vec3 {
	return Vec3{
		u[1]*v[2] - u[2]*v[1],
		u[2]*v[0] - u[0]*v[2],
		u[0]*v[1] - u[1]*v[0],
	}
}

func (u Vec3) UnitVector() Vec3 {
	return u.Div(u.Length())
}
