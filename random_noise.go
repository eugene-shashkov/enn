package enn

import (
	"math/rand"
)

func (n Network) RandomNoise(state *State) {
	layersLength := len(state.Layers)
	for i := 0; i < state.NoiseIterations; i++ {

		layer := rand.Intn(layersLength)

		if layer == 0 {
			layer++
		}

		weightLength := len(state.Layers[layer].Weights)
		neuronWeightKey := rand.Intn(weightLength)

		randNeuronWeightLength := len(state.Layers[layer].Weights[neuronWeightKey])
		randWeightKey := rand.Intn(randNeuronWeightLength)

		weight := state.Layers[layer].Weights[neuronWeightKey][randWeightKey]

		state.Layers[layer].Weights[neuronWeightKey][randWeightKey] = n.WeightGenerator(weight)
	}

}
