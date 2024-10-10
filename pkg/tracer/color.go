package tracer

import "fmt"

type Color = Vec3

func (c Color) ToPpm() string {
	intensity := Interval{Min: 0.000, Max: 0.999}
	r := int(256 * intensity.Clamp(c[0]))
	g := int(256 * intensity.Clamp(c[1]))
	b := int(256 * intensity.Clamp(c[2]))

	return fmt.Sprintf("%d %d %d", r, g, b)
}
