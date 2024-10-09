package tracer

type Ray struct {
	Origin    Point
	Direction Vector
}

func (ray Ray) At(t float64) Point {
	return ray.Origin.Add(ray.Direction.Mul(t))
}
