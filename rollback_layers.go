package enn

func (n Network) RollbackLayers(state *State) {
	// clear Layers
	state.Layers = []Layer{}
	for _, eachLayer := range state.RollbackLayers {

		var newlayer Layer

		for _, WeightGroup := range eachLayer.Weights {
			var newWeightGroup []float64
			for _, weight := range WeightGroup {
				newWeightGroup = append(newWeightGroup, weight)
			}

			newlayer.Weights = append(newlayer.Weights, newWeightGroup)
		}

		for _, neuron := range eachLayer.Neurons {
			newlayer.Neurons = append(newlayer.Neurons, neuron)
		}

		for _, bias := range eachLayer.Biases {
			newlayer.Biases = append(newlayer.Biases, bias)
		}

		state.Layers = append(state.Layers, newlayer)
	}
}
