package enn

import (
	"math"
)

func (n Network) Sigmoid(x float64) float64 {
	return 1 / (math.Abs(x) + 1)
}

func (n Network) SigmoidBack(y float64) float64 {
	return float64(1)/y - float64(1)
}

func (n Network) NumberToNeuron(x float64) float64 {

	// minus
	// from 0 to 0.5
	if x < 0 {
		return 0.5 / (1 + math.Abs(x))
	}

	// zero
	if x == 0 {
		return 0.5
	}

	// pluss
	// from 0.5 to 1
	return 1 - (0.5 / (1 + x))
}

func (n Network) NeuronToNumber(y float64) float64 {
	if y > 0 && y < 0.5 {
		return (0.5 - y) / -y
	}

	if y == 0.5 {
		return 0
	}

	return math.Abs(1 - 0.5/(1-y))
}
