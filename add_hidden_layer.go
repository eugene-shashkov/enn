package enn

import (
	"fmt"

	"github.com/the-power-of-code/enn/app"
)

func (n Network) AddHiddenLayer(NeuronsCount int, state *State) {
	// build hidden layer
	var layer Layer

	if len(state.Layers) == 0 {

		// getting length of input layer
		inputLayerNeuronsCount := len(state.Features[0])

		fmt.Println("inputLayerNeuronsCount", inputLayerNeuronsCount)

		layer.Weights = app.F64arr2(inputLayerNeuronsCount, 0, NeuronsCount)

	} else {

		// get length of last existed layer
		lastHiddenLayerNeuronCount := len(state.Layers[len(state.Layers)-1].Neurons)

		layer.Weights = app.F64arr2(lastHiddenLayerNeuronCount, 0, NeuronsCount)
	}

	// add neurons
	layer.Neurons = app.F64arr(NeuronsCount, 0.5)
	layer.Biases = app.F64arr(NeuronsCount, 0)

	state.Layers = append(state.Layers, layer)

}
