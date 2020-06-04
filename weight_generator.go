package enn

import (
	"math"
	"math/rand"

	"github.com/the-power-of-code/enn/app"
)

func (n Network) WeightGenerator(weight float64) float64 {
	sign := app.GoRand(1, 2)

	if sign == 1 {
		return weight - math.Abs(rand.NormFloat64())
	}

	return weight + math.Abs(rand.NormFloat64())
}
