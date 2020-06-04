package enn

import "github.com/the-power-of-code/enn/app"

func (n Network) SetDataset(features [][]float64, labels [][]float64, state *State) {
	state.Features = features
	state.Labels = labels
	state.FeaturesLength = app.GetDatasetLength(features)
	state.LabelsLength = app.GetDatasetLength(labels)

	state.Layers = append(state.Layers, n.NewInputLayer(features, state))

	n.DefineOutputNeurons(state)
}

func (n Network) NewInputLayer(features [][]float64, state *State) Layer {
	var inputlayer Layer

	// update input and output layer by new features
	for _, feature := range features[0] {
		var NeuronValue float64 = n.NumberToNeuron(feature)

		// fill input layer
		inputlayer.Neurons = append(inputlayer.Neurons, NeuronValue)
	}

	return inputlayer
}
