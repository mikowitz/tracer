package main

import (
	"math/rand/v2"
	"os"
	"runtime"
	"runtime/pprof"

	t "github.com/mikowitz/tracer/pkg/tracer"
)

func main() {
	file, _ := os.Create("./cpu.pprof")
	memfile, _ := os.Create("./mem.pprof")
	heapfile, _ := os.Create("./heap.pprof")
	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()
	defer pprof.Lookup("allocs").WriteTo(memfile, 0)
	defer pprof.Lookup("heap").WriteTo(heapfile, 0)
	defer runtime.GC()
	finalImage()
}

func finalImage() {
	world := t.HittableList{}

	groundMat := t.Lambertian{Albedo: t.Color{0.5, 0.5, 0.5}}
	world = append(world, &t.Sphere{Center: t.Point{0, -1000, 0}, Radius: 1000, Material: &groundMat})

	mat1 := t.Dielectric{RefractionIndex: 1.5}
	world = append(world, &t.Sphere{Center: t.Point{0, 1, 0}, Radius: 1.0, Material: &mat1})

	mat2 := t.Lambertian{Albedo: t.Color{0.4, 0.2, 0.1}}
	world = append(world, &t.Sphere{Center: t.Point{-4, 1, 0}, Radius: 1.0, Material: &mat2})

	mat3 := t.Metal{Albedo: t.Color{0.7, 0.6, 0.5}, Fuzz: 0}
	world = append(world, &t.Sphere{Center: t.Point{4, 1, 0}, Radius: 1.0, Material: &mat3})

	for a := range 22 {
		a -= 11
		for b := range 22 {
			b -= 11

			chooseMat := rand.Float64()
			center := t.Point{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}

			if center.Sub(t.Point{4, 0.2, 0}).Length() > 0.9 {
				if chooseMat < 0.8 {
					albedo := t.RandomVec().Prod(t.RandomVec())
					mat := t.Lambertian{Albedo: albedo}
					world = append(world, &t.Sphere{Center: center, Radius: 0.2, Material: &mat})
				} else if chooseMat < 0.95 {
					mat := t.Metal{Albedo: t.RandomVecIn(0.5, 1), Fuzz: t.RandomFloat64In(0, 0.5)}
					world = append(world, &t.Sphere{Center: center, Radius: 0.2, Material: &mat})
				} else {
					mat := t.Dielectric{RefractionIndex: 1.5}
					world = append(world, &t.Sphere{Center: center, Radius: 0.2, Material: &mat})
				}
			}
		}
	}

	camera := t.NewCamera(1200, 16.0/9.0)
	camera.SetSamplesPerPixel(100)
	camera.SetMaxDepth(10)

	camera.SetVerticalFieldOfView(20)
	camera.SetOrientation(
		t.Point{13, 2, 3},
		t.Point{0, 0, 0},
		t.Vector{0, 1, 0},
	)
	camera.SetFocus(0.6, 10)

	camera.Render(world)
}
