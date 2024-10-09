package main

import (
	"fmt"
	"os"
)

func main() {
	imageWidth := 256
	imageHeight := 256

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for y := range imageHeight {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d", imageHeight-y)
		for x := range imageWidth {
			r := 0.0
			g := float64(y) / float64(imageHeight-1)
			b := float64(x) / float64(imageWidth-1)

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
	fmt.Fprintf(os.Stderr, "\rDone.                     \n")
}
