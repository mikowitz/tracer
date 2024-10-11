package main

import (
	"math"

	t "github.com/mikowitz/tracer/pkg/tracer"
)

func main() {
	Vfov()
	// imageWidth := 400
	// aspectRatio := 16.0 / 9.0
	//
	// camera := t.NewCamera(imageWidth, aspectRatio)
	// camera.SetSamplesPerPixel(100)
	// camera.SetMaxDepth(50)
	// camera.SetVerticalFieldOfView(90)
	//
	// groundMat := t.Lambertian{Albedo: t.Color{0.8, 0.8, 0}}
	// centerMat := t.Lambertian{Albedo: t.Color{0.1, 0.2, 0.5}}
	// leftMat := t.Dielectric{RefractionIndex: 1.5}
	// bubbleMat := t.Dielectric{RefractionIndex: 1.00 / 1.5}
	// rightMat := t.Metal{Albedo: t.Color{0.8, 0.6, 0.2}, Fuzz: 0.8}
	//
	// world := t.HittableList{
	// 	&t.Sphere{Center: t.Point{0, -100.5, -1}, Radius: 100, Material: &groundMat},
	// 	&t.Sphere{Center: t.Point{0, 0, -1.2}, Radius: 0.5, Material: &centerMat},
	// 	&t.Sphere{Center: t.Point{-1, 0, -1}, Radius: 0.5, Material: &leftMat},
	// 	&t.Sphere{Center: t.Point{-1, 0, -1}, Radius: 0.4, Material: &bubbleMat},
	// 	&t.Sphere{Center: t.Point{1, 0, -1}, Radius: 0.5, Material: &rightMat},
	// }
	//
	// camera.Render(world)
}

func Vfov() {
	r := math.Cos(math.Pi / 4)

	leftMat := t.Lambertian{Albedo: t.Color{0, 0, 1}}
	rightMat := t.Lambertian{Albedo: t.Color{1, 0, 0}}

	world := t.HittableList{
		&t.Sphere{Center: t.Point{-r, 0, -1}, Radius: r, Material: &leftMat},
		&t.Sphere{Center: t.Point{r, 0, -1}, Radius: r, Material: &rightMat},
	}

	imageWidth := 400
	aspectRatio := 16.0 / 9.0

	camera := t.NewCamera(imageWidth, aspectRatio)
	camera.SetSamplesPerPixel(100)
	camera.SetMaxDepth(50)
	camera.SetVerticalFieldOfView(90)

	camera.Render(world)
}
