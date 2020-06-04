package main

import (
	"fmt"

	"github.com/the-power-of-code/enn"
)

// experimental neural network
// go run examples/basic_usage.go

func main() {

	var features [][]float64
	var labels [][]float64

	features = [][]float64{
		{10, 1},
		{100, 10},
	}

	labels = [][]float64{
		{1},
		{10},
	}

	nn := enn.CreateNeuralNetwork()

	state := nn.CreateState()

	// need to be declared before adding hidden layers
	nn.SetDataset(features, labels, &state)

	// use after SetDataset !

	nn.AddHiddenLayer(2, &state)
	nn.AddHiddenLayer(2, &state)

	nn.Learn(1000, &state)

	var testPrediction1 []float64 = nn.Predict(&state, features[0])
	var testPrediction2 []float64 = nn.Predict(&state, features[1])

	fmt.Println("Test Prediction Expected:", labels[0], testPrediction1)
	fmt.Println("Test Prediction Expected:", labels[1], testPrediction2)
	fmt.Println("final loss: ", state.Loss)

}
