package enn

type State struct {
	Features       [][]float64
	Labels         [][]float64
	FeaturesLength int64
	LabelsLength   int64

	Layers         []Layer
	RollbackLayers []Layer
	OutputNeurons  [][]float64

	Loss             float64
	MLInit           bool
	DatasetIteration int

	SuccessIterationCount int64
	FailIterationCount    int64

	NoiseIterations            int
	DeactivateNeuronIterations int
}

type Layer struct {
	Weights [][]float64 // for each neuron we will have collection of weights and one bias
	Neurons []float64
	Biases  []float64 // only one bias for each neuron
}

type Rollback struct {
	Layer  int
	Neuron int
	Weight int
	Value  float64
}
