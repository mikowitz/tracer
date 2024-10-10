package tracer

type Material interface {
	Scatter(ray Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool
}

type Lambertian struct {
	Albedo Color
}

type Metal struct {
	Albedo Color
	Fuzz   float64
}

type Dielectric struct {
	RefractionIndex float64
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
	reflected = reflected.UnitVector().Add(RandomUnitVector().Mul(m.Fuzz))
	*scattered = Ray{Origin: rec.P, Direction: reflected}
	*attenuation = m.Albedo
	return true
}

func (d *Dielectric) Scatter(ray Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	ri := d.RefractionIndex
	if rec.FrontFace {
		ri = 1.0 / d.RefractionIndex
	}

	unitDirection := ray.Direction.UnitVector()
	refracted := unitDirection.Refract(rec.Normal, ri)

	*scattered = Ray{Origin: rec.P, Direction: refracted}
	*attenuation = Color{1, 1, 1}

	return true
}
