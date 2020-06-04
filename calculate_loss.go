package enn

import (
	"math"
)

func (n Network) CalculateLoss(state *State) {

	var loss float64
	var lastlayerkey int = len(state.Layers) - 1

	for key, eachNeuronHave := range state.OutputNeurons[state.DatasetIteration] {

		loss = loss + math.Abs(eachNeuronHave-state.Layers[lastlayerkey].Neurons[key])
	}

	if state.MLInit == false {
		state.MLInit = true
		state.Loss = loss
		return
	}

	if loss > state.Loss {
		n.RollbackLayers(state)

		state.FailIterationCount++
		return
	}

	state.SuccessIterationCount++
	state.Loss = loss

}
