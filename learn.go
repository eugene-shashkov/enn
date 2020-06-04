package enn

import (
	"github.com/the-power-of-code/enn/app"
)

func (n Network) Learn(epochs int64, state *State) {

	state.Layers = append(state.Layers, n.NewOutputLayer(state))

	for ep := int64(0); ep < epochs; ep++ {
		n.LearnWorker(state)

	}

}

func (n Network) NewOutputLayer(state *State) Layer {

	var outputLayer Layer

	outputLayer.Weights = app.F64arr2(len(state.Layers[len(state.Layers)-1].Neurons), 0.1, len(state.Labels[0]))
	outputLayer.Neurons = app.F64arr(len(state.Labels[0]), 0)
	outputLayer.Biases = app.F64arr(len(state.Labels[0]), 0)

	return outputLayer
}
