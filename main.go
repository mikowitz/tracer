package main

import (
	"fmt"
	"os"

	t "github.com/mikowitz/tracer/pkg/tracer"
)

func main() {
	imageWidth := 256
	imageHeight := 256

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for y := range imageHeight {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d", imageHeight-y)
		for x := range imageWidth {
			color := t.Color{0, float64(y) / float64(imageHeight-1), float64(x) / float64(imageWidth-1)}
			fmt.Println(color.ToPpm())
		}
	}
	fmt.Fprintf(os.Stderr, "\rDone.                     \n")
}
