package enn

func (n Network) Predict(state *State, set []float64) []float64 {

	state.Layers[0] = n.FillInputPredict(state, set)

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

	var labels []float64

	for _, eachNeuron := range state.Layers[len(state.Layers)-1].Neurons {
		labels = append(labels, n.NeuronToNumber(eachNeuron))
	}

	return labels
}

func (n Network) FillInputPredict(state *State, features []float64) Layer {
	var inputlayer Layer

	// update input and output layer by new features
	for _, feature := range features {
		var NeuronValue float64 = n.NumberToNeuron(feature)

		// fill input layer
		inputlayer.Neurons = append(inputlayer.Neurons, NeuronValue)
	}

	return inputlayer
}
