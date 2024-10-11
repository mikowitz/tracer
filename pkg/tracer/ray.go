package tracer

type Ray struct {
	Origin    Point
	Direction Vector
	Time      float64
}

func (ray Ray) At(t float64) Point {
	return ray.Origin.Add(ray.Direction.Mul(t))
}
