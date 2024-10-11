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

	vfov             float64
	lookfrom, lookat Point
	vup              Vector
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

func (c *Camera) SetVerticalFieldOfView(vfov float64) {
	c.vfov = vfov
}

func (c *Camera) SetOrientation(lookfrom, lookat, vup Vec3) {
	c.lookfrom = lookfrom
	c.lookat = lookat
	c.vup = vup
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

	c.center = c.lookfrom

	focalLength := c.lookfrom.Sub(c.lookat).Length()
	θ := DegreesToRadians(c.vfov)
	h := math.Tan(θ / 2)
	viewportHeight := 2 * h * focalLength
	viewportWidth := viewportHeight * (float64(c.imageWidth) / float64(c.imageHeight))

	w := c.lookfrom.Sub(c.lookat).UnitVector()
	u := c.vup.Cross(w).UnitVector()
	v := w.Cross(u)

	viewportU := u.Mul(viewportWidth)
	viewportV := v.Neg().Mul(viewportHeight)

	c.pixelΔU = viewportU.Div(float64(c.imageWidth))
	c.pixelΔV = viewportV.Div(float64(c.imageHeight))

	viewportUpperLeft := c.center.Sub(w.Mul(focalLength)).Sub(viewportU.Div(2)).Sub(viewportV.Div(2))
	c.pixel00Loc = viewportUpperLeft.Add(c.pixelΔU.Add(c.pixelΔV).Mul(0.5))
}

func (c Camera) rayColor(ray Ray, world HittableList, depth int) Color {
	if depth <= 0 {
		return Color{0, 0, 0}
	}

	rec := HitRecord{}

	if world.Hit(ray, Interval{Min: 0.001, Max: math.Inf(1)}, &rec) {
		scattered := Ray{}
		attenuation := Color{}
		if rec.Material.Scatter(ray, rec, &attenuation, &scattered) {
			return c.rayColor(scattered, world, depth-1).Prod(attenuation)
		}
		return Color{0, 0, 0}
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
