package enn

type Actions interface {
	Learn(epochs int64, state *State)
	CreateState() State

	AddHiddenLayer(NeuronsCount int, state *State)
	NewInputLayer(features [][]float64, state *State) Layer
	NewOutputLayer(state *State) Layer

	SetDataset(features [][]float64, labels [][]float64, state *State)
	LearnWorker(state *State)

	Sigmoid(float64) float64
	SigmoidBack(float64) float64

	NumberToNeuron(float64) float64
	NeuronToNumber(float64) float64

	WeightGenerator(float64) float64

	RecalculateState(state *State)
	CalculateLoss(state *State)

	DefineOutputNeurons(state *State)

	FillInputPredict(state *State, features []float64) Layer
	Predict(state *State, set []float64) []float64

	BackupLayers(state *State)
	RollbackLayers(state *State)
}

type Network struct {
}

func CreateNeuralNetwork() Actions {
	return &Network{}
}
