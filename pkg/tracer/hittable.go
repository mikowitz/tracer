package tracer

type Hittable interface {
	Hit(ray Ray, min, max float64, rec *HitRecord) bool
}
