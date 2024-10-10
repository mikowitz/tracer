package tracer

import (
	"fmt"
	"math"
)

type Color = Vec3

func (c Color) ToPpm() string {
	intensity := Interval{Min: 0.000, Max: 0.999}
	r := int(256 * intensity.Clamp(linearToGamma(c[0])))
	g := int(256 * intensity.Clamp(linearToGamma(c[1])))
	b := int(256 * intensity.Clamp(linearToGamma(c[2])))

	return fmt.Sprintf("%d %d %d", r, g, b)
}

func linearToGamma(linear float64) float64 {
	if linear > 0.0 {
		return math.Sqrt(linear)
	}
	return 0
}
