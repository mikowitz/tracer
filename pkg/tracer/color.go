package tracer

import "fmt"

type Color = Vec3

func (c Color) ToPpm() string {
	r := int(255.999 * c[0])
	g := int(255.999 * c[1])
	b := int(255.999 * c[2])

	return fmt.Sprintf("%d %d %d", r, g, b)
}
