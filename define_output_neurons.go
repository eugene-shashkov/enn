package enn

func (n Network) DefineOutputNeurons(state *State) {
	for _, eachFeatureRow := range state.Labels {
		var neurons []float64
		for _, eachFeature := range eachFeatureRow {
			var neuron float64 = n.NumberToNeuron(eachFeature)
			neurons = append(neurons, neuron)
		}
		state.OutputNeurons = append(state.OutputNeurons, neurons)
	}
}
