package main

import (
	"fmt"
	"os"

	t "github.com/mikowitz/tracer/pkg/tracer"
)

func main() {
	imageWidth := 400
	aspectRatio := 16.0 / 9.0

	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}

	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := t.Point{0, 0, 0}

	viewportU := t.Vector{viewportWidth, 0, 0}
	viewportV := t.Vector{0, -viewportHeight, 0}

	pixelDeltaU := viewportU.Div(float64(imageWidth))
	pixelDeltaV := viewportV.Div(float64(imageHeight))

	viewportUpperLeft := cameraCenter.Sub(t.Vector{0, 0, focalLength}).Sub(viewportU.Div(2)).Sub(viewportV.Div(2))
	pixel00Loc := viewportUpperLeft.Add(pixelDeltaU.Add(pixelDeltaV).Mul(0.5))

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for y := range imageHeight {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d", imageHeight-y)
		for x := range imageWidth {
			pixelCenter := pixel00Loc.Add(pixelDeltaU.Mul(float64(x))).Add(pixelDeltaV.Mul(float64(y)))
			rayDirection := pixelCenter.Sub(cameraCenter)
			ray := t.Ray{Origin: cameraCenter, Direction: rayDirection}
			color := RayColor(ray)
			fmt.Println(color.ToPpm())
		}
	}
	fmt.Fprintf(os.Stderr, "\rDone.                     \n")
}

func RayColor(r t.Ray) t.Color {
	if HitSphere(t.Point{0, 0, -1}, 0.5, r) {
		return t.Color{1, 0, 0}
	}
	unitDirection := r.Direction.UnitVector()
	a := 0.5 * (unitDirection[1] + 1.0)
	return t.Color{1.0, 1.0, 1.0}.Mul(1.0 - a).Add(t.Color{0.5, 0.7, 1.0}.Mul(a))
}

func HitSphere(center t.Point, radius float64, ray t.Ray) bool {
	oc := center.Sub(ray.Origin)
	a := ray.Direction.Dot(ray.Direction)
	b := -2.0 * ray.Direction.Dot(oc)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c

	return discriminant >= 0
}
