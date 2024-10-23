package main

import (
	"enn"
	"fmt"
)

// experimental neural network
// go run examples/basic_usage.go

func main() {
	var features [][]float64
	var labels [][]float64

	features = [][]float64{
		{1, 1, 3, 1, 2},
		{15, 13, 10, 9, 13},
		{99, 100, 92, 101, 119},
	}

	labels = [][]float64{
		{2},
		{30},
		{200},
	}

	nn := enn.CreateNeuralNetwork()

	state := nn.CreateState()

	// need to be declared before adding hidden layers
	nn.SetDataset(features, labels, &state)

	// use after SetDataset !

	nn.AddHiddenLayer(5, &state)

	nn.Learn(100, &state)

	var testPrediction1 []float64 = nn.Predict(&state, features[0])
	var testPrediction2 []float64 = nn.Predict(&state, features[1])
	var testPrediction3 []float64 = nn.Predict(&state, features[2])

	fmt.Println("Test Prediction Expected:", labels[0], testPrediction1)
	fmt.Println("Test Prediction Expected:", labels[1], testPrediction2)
	fmt.Println("Test Prediction Expected:", labels[2], testPrediction3)
	fmt.Println("final loss: ", state.Loss)

}
