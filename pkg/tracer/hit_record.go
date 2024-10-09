package tracer

type HitRecord struct {
	P      Point
	Normal Vector
	T      float64
}

type Hittable interface {
	Hit(ray Ray, min, max float64, rec *HitRecord) bool
}
