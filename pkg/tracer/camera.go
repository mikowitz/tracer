package tracer

import (
	"fmt"
	"math"
	"os"
)

type Camera struct {
	imageWidth, imageHeight int
	aspectRatio             float64
	center                  Point
	pixelΔU, pixelΔV        Vector
	pixel00Loc              Point
}

func NewCamera(imageWidth int, aspectRatio float64) Camera {
	return Camera{
		imageWidth:  imageWidth,
		aspectRatio: aspectRatio,
	}
}

func (c *Camera) Render(world HittableList) {
	c.initialize()
	fmt.Printf("P3\n%d %d\n255\n", c.imageWidth, c.imageHeight)

	for y := range c.imageHeight {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d", c.imageHeight-y)
		for x := range c.imageWidth {
			pixelCenter := c.pixel00Loc.Add(c.pixelΔU.Mul(float64(x))).Add(c.pixelΔV.Mul(float64(y)))
			rayDirection := pixelCenter.Sub(c.center)
			ray := Ray{Origin: c.center, Direction: rayDirection}
			color := c.rayColor(ray, world)
			fmt.Println(color.ToPpm())
		}
	}
	fmt.Fprintf(os.Stderr, "\rDone.                     \n")
}

func (c *Camera) initialize() {
	c.imageHeight = int(float64(c.imageWidth) / c.aspectRatio)
	if c.imageHeight < 1 {
		c.imageHeight = 1
	}

	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(c.imageWidth) / float64(c.imageHeight))
	c.center = Point{0, 0, 0}

	viewportU := Vector{viewportWidth, 0, 0}
	viewportV := Vector{0, -viewportHeight, 0}

	c.pixelΔU = viewportU.Div(float64(c.imageWidth))
	c.pixelΔV = viewportV.Div(float64(c.imageHeight))

	viewportUpperLeft := c.center.Sub(Vector{0, 0, focalLength}).Sub(viewportU.Div(2)).Sub(viewportV.Div(2))
	c.pixel00Loc = viewportUpperLeft.Add(c.pixelΔU.Add(c.pixelΔV).Mul(0.5))
}

func (c Camera) rayColor(ray Ray, world HittableList) Color {
	rec := HitRecord{}

	if world.Hit(ray, Interval{Min: 0, Max: math.Inf(1)}, &rec) {
		return rec.Normal.Add(Color{1, 1, 1}).Mul(0.5)
	}

	unitDirection := ray.Direction.UnitVector()
	a := 0.5 * (unitDirection[1] + 1.0)
	return Color{1.0, 1.0, 1.0}.Mul(1.0 - a).Add(Color{0.5, 0.7, 1.0}.Mul(a))
}
