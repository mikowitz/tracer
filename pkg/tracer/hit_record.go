package tracer

type HitRecord struct {
	P         Point
	Normal    Vector
	T         float64
	FrontFace bool
}

func (hr *HitRecord) SetFaceNormal(ray Ray, outwardNormal Vector) {
	frontFace := ray.Direction.Dot(outwardNormal) < 0
	if frontFace {
		hr.Normal = outwardNormal
	} else {
		hr.Normal = outwardNormal.Neg()
	}

}
