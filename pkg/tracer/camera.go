package tracer

import (
	"fmt"
	"math"
	"math/rand/v2"
	"os"
)

type Camera struct {
	imageWidth, imageHeight int
	aspectRatio             float64
	center                  Point
	pixelΔU, pixelΔV        Vector
	pixel00Loc              Point

	samplesPerPixel   int
	pixelsSampleScale float64
	maxDepth          int
}

func NewCamera(imageWidth int, aspectRatio float64) Camera {
	return Camera{
		imageWidth:  imageWidth,
		aspectRatio: aspectRatio,
	}
}

func (c *Camera) SetSamplesPerPixel(samples int) {
	c.samplesPerPixel = samples
}

func (c *Camera) SetMaxDepth(depth int) {
	c.maxDepth = depth
}

func (c *Camera) Render(world HittableList) {
	c.initialize()
	fmt.Printf("P3\n%d %d\n255\n", c.imageWidth, c.imageHeight)

	for y := range c.imageHeight {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %5d", c.imageHeight-y)
		for x := range c.imageWidth {
			color := Color{0, 0, 0}
			for _ = range c.samplesPerPixel {
				ray := c.getRay(x, y)
				color = color.Add(c.rayColor(ray, world, c.maxDepth))
			}
			fmt.Println(color.Mul(c.pixelsSampleScale).ToPpm())
		}
	}
	fmt.Fprintf(os.Stderr, "\rDone.                     \n")
}

func (c *Camera) initialize() {
	c.imageHeight = int(float64(c.imageWidth) / c.aspectRatio)
	if c.imageHeight < 1 {
		c.imageHeight = 1
	}

	c.pixelsSampleScale = 1.0 / float64(c.samplesPerPixel)

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

func (c Camera) rayColor(ray Ray, world HittableList, depth int) Color {
	if depth <= 0 {
		return Color{0, 0, 0}
	}

	rec := HitRecord{}

	if world.Hit(ray, Interval{Min: 0.001, Max: math.Inf(1)}, &rec) {
		direction := RandomOnHemisphere(rec.Normal)
		return c.rayColor(Ray{Origin: rec.P, Direction: direction}, world, depth-1).Mul(0.5)
	}

	unitDirection := ray.Direction.UnitVector()
	a := 0.5 * (unitDirection[1] + 1.0)
	return Color{1.0, 1.0, 1.0}.Mul(1.0 - a).Add(Color{0.5, 0.7, 1.0}.Mul(a))
}

func (c Camera) getRay(x, y int) Ray {
	xOffset := rand.Float64() - 0.5
	yOffset := rand.Float64() - 0.5

	pixelSample := c.pixel00Loc.Add(c.pixelΔU.Mul(xOffset + float64(x))).Add(c.pixelΔV.Mul(yOffset + float64(y)))

	direction := pixelSample.Sub(c.center)
	return Ray{Origin: c.center, Direction: direction}
}
