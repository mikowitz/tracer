package tracer

import (
	"math"
	"math/rand/v2"
)

type Material interface {
	Scatter(ray Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool
}

type Lambertian struct {
	Albedo Color
}

type Metal struct {
	Albedo Color
	Fuzz   float64
}

type Dielectric struct {
	RefractionIndex float64
}

func (l *Lambertian) Scatter(ray Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	scatterDirection := rec.Normal.Add(RandomUnitVector())
	if scatterDirection.IsNearZero() {
		scatterDirection = rec.Normal
	}
	*scattered = Ray{Origin: rec.P, Direction: scatterDirection, Time: ray.Time}
	*attenuation = l.Albedo
	return true
}

func (m *Metal) Scatter(ray Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	reflected := ray.Direction.Reflect(rec.Normal)
	reflected = reflected.UnitVector().Add(RandomUnitVector().Mul(m.Fuzz))
	*scattered = Ray{Origin: rec.P, Direction: reflected, Time: ray.Time}
	*attenuation = m.Albedo
	return true
}

func (d *Dielectric) Scatter(ray Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	ri := d.RefractionIndex
	if rec.FrontFace {
		ri = 1.0 / d.RefractionIndex
	}

	unitDirection := ray.Direction.UnitVector()
	cosθ := math.Min(unitDirection.Neg().Dot(rec.Normal), 1.0)
	sinθ := math.Sqrt(1.0 - cosθ*cosθ)
	cannotRefract := ri*sinθ > 1.0

	direction := unitDirection.Refract(rec.Normal, ri)
	if cannotRefract || reflectance(cosθ, ri) > rand.Float64() {
		direction = unitDirection.Reflect(rec.Normal)
	}

	*scattered = Ray{Origin: rec.P, Direction: direction, Time: ray.Time}
	*attenuation = Color{1, 1, 1}

	return true
}

func reflectance(cosine, refractionIndex float64) float64 {
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 *= r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
