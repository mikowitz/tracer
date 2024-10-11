package tracer

type Hittable interface {
	Hit(ray Ray, interval Interval, rec *HitRecord) bool
	BoundingBox() Aabb
}
