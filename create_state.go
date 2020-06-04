package enn

func (n Network) CreateState() State {
	return State{
		NoiseIterations:            15,
		DeactivateNeuronIterations: 3,
	}
}
