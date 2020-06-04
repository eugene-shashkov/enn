package enn

func (n Network) LearnWorker(state *State) {

	var lastLayerKey int = len(state.Layers) - 1

	// update input layer by new features
	for key, featureRow := range state.Features {

		// set feature iteration key
		state.DatasetIteration = key

		// clear input and output layers
		state.Layers[0].Neurons = []float64{}
		state.Layers[lastLayerKey].Neurons = []float64{}

		for _, feature := range featureRow {
			var inputNeuron float64 = n.NumberToNeuron(feature)

			// fill input layer
			state.Layers[0].Neurons = append(state.Layers[0].Neurons, inputNeuron)
		}

		// fill output layer
		for _, label := range state.Labels[key] {
			var outputNeuron float64 = n.NumberToNeuron(label)

			state.Layers[lastLayerKey].Neurons = append(state.Layers[lastLayerKey].Neurons, outputNeuron)
		}

		n.BackupLayers(state)
		n.RandomNoise(state)

		n.RecalculateState(state)
		n.CalculateLoss(state)

		n.DeactivateNeuron(state)

	}

}
