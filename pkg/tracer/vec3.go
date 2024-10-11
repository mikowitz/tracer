package tracer

import (
	"math"
	"math/rand/v2"
)

type Vec3 [3]float64
type Vector = Vec3
type Point = Vec3

func RandomVec() Vec3 {
	return Vec3{
		rand.Float64(),
		rand.Float64(),
		rand.Float64(),
	}
}

func RandomVecIn(min, max float64) Vec3 {
	return Vec3{
		RandomFloat64In(min, max),
		RandomFloat64In(min, max),
		RandomFloat64In(min, max),
	}
}

func RandomUnitVector() Vec3 {
	for {
		p := RandomVecIn(-1, 1)
		lensq := p.LengthSquared()
		if 1e-160 < lensq && lensq <= 1.0 {
			return p.UnitVector()
		}
	}
}

func RandomOnHemisphere(normal Vector) Vec3 {
	onUnitSphere := RandomUnitVector()
	if onUnitSphere.Dot(normal) > 0.0 {
		return onUnitSphere
	}
	return onUnitSphere.Neg()
}

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

func (u Vec3) Reflect(n Vec3) Vector {
	return u.Sub(n.Mul(2.0 * u.Dot(n)))
}

func (u Vec3) Refract(n Vec3, ηiOverηt float64) Vec3 {
	cosθ := math.Min(u.Neg().Dot(n), 1.0)
	rOutPerp := u.Add(n.Mul(cosθ)).Mul(ηiOverηt)
	rOutParallel := n.Mul(-math.Sqrt(math.Abs(1.0 - rOutPerp.LengthSquared())))
	return rOutPerp.Add(rOutParallel)
}

func (u Vec3) UnitVector() Vec3 {
	return u.Div(u.Length())
}

func (u Vec3) IsNearZero() bool {
	s := 1e-8
	return math.Abs(u[0]) < s && math.Abs(u[1]) < s && math.Abs(u[2]) < s
}
