package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/stat"
)

func TestGenerateFloatData(t *testing.T) {
	// Arrange
	mean := 50.0
	std := 10.0
	size := 1000

	// Act
	data := GenerateFloatData(mean, std, size)

	// Assert
	calculatedMean := stat.Mean(data, nil)
	calculatedStddev := stat.StdDev(data, nil)
	assert.InDelta(t, mean, calculatedMean, 0.5)
	assert.InDelta(t, std, calculatedStddev, 0.5)
	assert.Len(t, data, size)

}

func TestGenerateIntegerData(t *testing.T) {
	// Arrange
	mean := 50.0
	std := 10.0
	size := 1000

	// Act
	data := GenerateIntegerData(mean, std, size)

	// Assert
	assert.InDelta(t, mean, calculateAverage(data, size), 0.5)
	assert.Len(t, data, size)
}

func TestGenerateStringData(t *testing.T) {
	// Arrange
	values := []string{"a", "b", "c"}
	size := 1000

	// Act
	data := GenerateStringData(values, size)

	// Assert
	assert.Len(t, data, size)

}

func calculateAverage(data []int64, size int) float64 {
	var sum int64

	for _, v := range data {
		sum += v
	}

	return float64(sum) / float64(size)
}
