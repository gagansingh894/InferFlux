package generator

import (
	"math/rand"

	"github.com/gagansingh894/InferFlux/types"
)

type DataGen map[string]interface{}

type Generator interface {
	GenerateWithRandomSize(modelSpec types.ModelSpec) DataGen
	GenerateWithFixedSize(modelSpec types.ModelSpec, size int) (DataGen, error)
	GenerateWithinRange(modelSpec types.ModelSpec, lowerBound int, higherBound int) (DataGen, error)
}

func GenerateStringData(values []string, size int) []string {
	out := make([]string, size)

	for i := range out {
		out[i] = values[rand.Intn(size)]
	}

	return out
}

func GenerateFloatData(mean float64, std float64, size int) []float64 {
	out := make([]float64, size)

	for i := range out {
		out[i] = rand.NormFloat64()*std + mean
	}

	return out
}

func GenerateIntegerData(mean float64, std float64, size int) []int64 {
	out := make([]int64, size)

	for i := range out {
		out[i] = int64(rand.NormFloat64()*std + mean)
	}

	return out
}
