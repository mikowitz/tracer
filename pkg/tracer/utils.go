package tracer

import (
	"math"
	"math/rand/v2"
)

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func RandomIn(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}
