package tracer

type HittableList []Hittable

func (hl *HittableList) Hit(ray Ray, interval Interval, rec *HitRecord) bool {
	tempRec := HitRecord{}
	hitAnything := false
	closestSoFar := interval.Max

	for _, object := range *hl {
		if object.Hit(ray, Interval{Min: interval.Min, Max: closestSoFar}, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.T
			*rec = tempRec
		}
	}

	return hitAnything
}

func (hl *HittableList) BoundingBox() Aabb {
	var bbox Aabb
	for _, object := range *hl {
		bbox = NewBoundingBoxFromBoundingBoxes(bbox, object.BoundingBox())
	}
	return bbox
}
