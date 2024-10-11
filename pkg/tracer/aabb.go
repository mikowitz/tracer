package tracer

type Aabb struct {
	X, Y, Z Interval
}

func NewBoundingBox(x, y, z Interval) Aabb {
	return Aabb{X: x, Y: y, Z: z}
}

func NewBoundingBoxFromPoints(a, b Point) Aabb {
	x := NewInterval(a[0], b[0])
	if a[0] <= b[0] {
		x = NewInterval(b[0], a[0])
	}
	y := NewInterval(a[1], b[1])
	if a[1] <= b[1] {
		y = NewInterval(b[1], a[1])
	}
	z := NewInterval(a[2], b[2])
	if a[2] <= b[2] {
		z = NewInterval(b[2], a[2])
	}
	return NewBoundingBox(x, y, z)
}

func (aabb Aabb) AxisInterval(n int) Interval {
	if n == 1 {
		return aabb.Y
	}
	if n == 2 {
		return aabb.Z
	}
	return aabb.X
}

func (aabb *Aabb) Hit(ray Ray, interval Interval, rec *HitRecord) bool {
	rayOrig := ray.Origin
	rayDir := ray.Direction

	for axis := range 3 {
		ax := aabb.AxisInterval(axis)
		adinv := 1.0 / rayDir[axis]

		t0 := (ax.Min - rayOrig[axis]) * adinv
		t1 := (ax.Max - rayOrig[axis]) * adinv

		if t0 < t1 {
			if t0 > interval.Min {
				interval.Min = t0
			}
			if t1 < interval.Max {
				interval.Max = t1
			}
		} else {
			if t1 > interval.Min {
				interval.Min = t1
			}
			if t0 < interval.Max {
				interval.Max = t0
			}
		}

		if interval.Max <= interval.Min {
			return false
		}
	}
	return true
}
