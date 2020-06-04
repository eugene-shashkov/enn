package enn

func (n Network) DeactivateNeuron(state *State) {

	for layerKey, layer := range state.Layers {

		for weightsKey, weights := range layer.Weights {

			for weightKey, _ := range weights {
				n.BackupLayers(state)
				state.Layers[layerKey].Weights[weightsKey][weightKey] = 0
				n.RecalculateState(state)
				n.CalculateLoss(state)
			}

		}

	}

}
