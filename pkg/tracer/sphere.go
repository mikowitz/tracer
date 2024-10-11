package tracer

import "math"

type Sphere struct {
	Center   Ray
	Radius   float64
	Material Material
	bbox     Aabb
}

func NewSphere(center Point, radius float64, mat Material) *Sphere {
	rvec := Vector{radius, radius, radius}
	return &Sphere{
		Center:   Ray{Origin: center, Direction: Vector{0, 0, 0}},
		Radius:   math.Max(0, radius),
		Material: mat,
		bbox:     NewBoundingBoxFromPoints(center.Sub(rvec), center.Add(rvec)),
	}
}

func MovingSphere(center1, center2 Point, radius float64, mat Material) *Sphere {
	rvec := Vector{radius, radius, radius}
	box1 := NewBoundingBoxFromPoints(center1.Sub(rvec), center1.Add(rvec))
	box2 := NewBoundingBoxFromPoints(center2.Sub(rvec), center2.Add(rvec))
	return &Sphere{
		Center:   Ray{Origin: center1, Direction: center2.Sub(center1)},
		Radius:   math.Max(0, radius),
		Material: mat,
		bbox:     NewBoundingBoxFromBoundingBoxes(box1, box2),
	}
}

func (s *Sphere) Hit(ray Ray, interval Interval, rec *HitRecord) bool {
	center := s.Center.At(ray.Time)
	oc := center.Sub(ray.Origin)
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
	outwardNormal := rec.P.Sub(center).Div(s.Radius)
	(*rec).SetFaceNormal(ray, outwardNormal)
	(*rec).Material = s.Material

	return true
}

func (s *Sphere) BoundingBox() Aabb {
	return s.bbox
}
