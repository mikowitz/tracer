package tracer

type BvhNode struct {
	left, right Hittable
	bbox        Aabb
}

func (bvh *BvhNode) Hit(ray Ray, interval Interval, rec *HitRecord) bool {
	if !bvh.bbox.Hit(ray, interval, rec) {
		return false
	}

	hitLeft := bvh.left.Hit(ray, interval, rec)
	hitRightInterval := NewInterval(interval.Min, interval.Max)
	if hitLeft {
		hitRightInterval.Max = rec.T
	}
	hitRight := bvh.right.Hit(ray, hitRightInterval, rec)

	return hitLeft || hitRight
}

func (bvh *BvhNode) BoundingBox() Aabb {
	return bvh.bbox
}
