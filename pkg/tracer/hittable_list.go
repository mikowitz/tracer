package tracer

type HittableList []Hittable

func (hl *HittableList) Hit(ray Ray, min, max float64, rec *HitRecord) bool {
	tempRec := HitRecord{}
	hitAnything := false
	closestSoFar := max

	for _, object := range *hl {
		if object.Hit(ray, min, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.T
			*rec = tempRec
		}
	}

	return hitAnything
}
