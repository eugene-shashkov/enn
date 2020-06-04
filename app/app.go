package app

import (
	"math/rand"
	"time"
)

func GoRand(min int64, max int64) int64 {
	max++

	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}

func GoRandInt(min int, max int) int {
	max++
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min

}

func GetDatasetLength(F [][]float64) int64 {
	var length int64
	for range F {
		length++
	}
	return length
}

func GetLength(L []float64) int64 {
	var length int64
	for range L {
		length++
	}
	return length
}

func GetEquationlength(equation [][]int64) int64 {
	var i int64
	for range equation {
		i++
	}
	return i
}

func GetOperandsLength(L []int64) int64 {
	var length int64
	for range L {
		length++
	}
	return length
}

func Reverse(arr []int64) []int64 {

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func Ftoi(f64 float64) int64 {
	var i64 int64
	for i := float64(1); i <= f64; i++ {
		i64++
	}
	return i64
}

func F64arr(length int, values float64) []float64 {

	var arr []float64

	for i := 0; i < length; i++ {
		arr = append(arr, values)
	}

	return arr
}

func F64arr2(length int, values float64, count int) [][]float64 {
	var arr2 [][]float64

	for c := 0; c < count; c++ {

		var arr []float64
		for i := 0; i < length; i++ {
			arr = append(arr, values)
		}

		arr2 = append(arr2, arr)
	}

	return arr2
}

func F64arr3(length int, values float64, count int, count2 int) [][][]float64 {

	var arr3 [][][]float64
	for x := 0; x < count2; x++ {
		var arr2 [][]float64
		for c := 0; c < count; c++ {
			var arr []float64
			for i := 0; i < length; i++ {
				arr = append(arr, values)
			}
			arr2 = append(arr2, arr)
		}
		arr3 = append(arr3, arr2)
	}

	return arr3
}
