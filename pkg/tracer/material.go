package tracer

type Material interface {
	Scatter(ray Ray, rec *HitRecord, attenuation *Color, scattered *Ray) bool
}
