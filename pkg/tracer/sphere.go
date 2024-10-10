package tracer

import "math"

type Sphere struct {
	Center   Point
	Radius   float64
	Material Material
}

func (s *Sphere) Hit(ray Ray, interval Interval, rec *HitRecord) bool {
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
	if !interval.Contains(root) {
		root = (h + sqrtd) / a
		if !interval.Contains(root) {
			return false
		}
	}

	(*rec).T = root
	(*rec).P = ray.At(rec.T)
	outwardNormal := rec.P.Sub(s.Center).Div(s.Radius)
	(*rec).SetFaceNormal(ray, outwardNormal)
	(*rec).Material = s.Material

	return true
}
