package tracer

type Material interface {
	Scatter(ray Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool
}

type Lambertian struct {
	Albedo Color
}

type Metal struct {
	Albedo Color
}

func (l *Lambertian) Scatter(ray Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	scatterDirection := rec.Normal.Add(RandomUnitVector())
	if scatterDirection.IsNearZero() {
		scatterDirection = rec.Normal
	}
	*scattered = Ray{Origin: rec.P, Direction: scatterDirection}
	*attenuation = l.Albedo
	return true
}

func (m *Metal) Scatter(ray Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	reflected := ray.Direction.Reflect(rec.Normal)
	*scattered = Ray{Origin: rec.P, Direction: reflected}
	*attenuation = m.Albedo
	return true
}
