package main

import (
	t "github.com/mikowitz/tracer/pkg/tracer"
)

func main() {
	imageWidth := 400
	aspectRatio := 16.0 / 9.0

	camera := t.NewCamera(imageWidth, aspectRatio)
	camera.SetSamplesPerPixel(100)
	camera.SetMaxDepth(50)

	world := t.HittableList{
		&t.Sphere{Center: t.Point{0, 0, -1}, Radius: 0.5},
		&t.Sphere{Center: t.Point{0, -100.5, -1}, Radius: 100},
	}

	camera.Render(world)
}
