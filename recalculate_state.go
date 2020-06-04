package enn

func (n Network) RecalculateState(state *State) {

	for layerKey, layer := range state.Layers {

		if layerKey == 0 {
			continue
		}
		var previousNeuron float64

		for weightNeuronKey, weights := range layer.Weights {

			var neuron float64

			for weightKey, weightValue := range weights {
				previousNeuron = state.Layers[layerKey-1].Neurons[weightKey]

				neuron = neuron + previousNeuron*weightValue
			}
			sNeuron := n.Sigmoid(neuron)
			state.Layers[layerKey].Neurons[weightNeuronKey] = sNeuron

		}
	}

}
