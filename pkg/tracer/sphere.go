package tracer

import "math"

type Sphere struct {
	Center Point
	Radius float64
}

func (s *Sphere) Hit(ray Ray, min, max float64, rec *HitRecord) bool {
	oc := s.Center.Sub(ray.Origin)
	a := ray.Direction.LengthSquared()
	h := ray.Direction.Dot(oc)
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := h*h - a*c
	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)

	root := (h - sqrtd) / a
	if root <= min || max <= root {
		root = (h + sqrtd) / a
		if root <= min || max <= root {
			return false
		}
	}

	rec.T = root
	rec.P = ray.At(rec.T)
	outwardNormal := rec.P.Sub(s.Center).Div(s.Radius)
	rec.SetFaceNormal(ray, outwardNormal)

	return true
}
