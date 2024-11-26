package agents

import (
	"errors"
	"math/rand"

	"github.com/gagansingh894/InferFlux/pkg/generator"
	"github.com/gagansingh894/InferFlux/types"
)

type HTTPAgent struct{}

func NewHTTPAgent() (*HTTPAgent, error) {
	return &HTTPAgent{}, nil
}

func (a *HTTPAgent) GenerateWithRandomSize(modelSpec types.ModelSpec) generator.DataGen {
	out := make(generator.DataGen)
	size := rand.Intn(1500) + 1

	for key, value := range modelSpec {
		switch value.Dtype {
		case types.String:
			out[key] = generator.GenerateStringData(value.Constraint.StringConstraint.Values, size)
		case types.Int:
			out[key] = generator.GenerateIntegerData(
				value.Constraint.NumericConstraint.Mean,
				value.Constraint.NumericConstraint.StandardDeviation,
				size,
			)
		case types.Float:
			out[key] = generator.GenerateFloatData(
				value.Constraint.NumericConstraint.Mean,
				value.Constraint.NumericConstraint.StandardDeviation,
				size,
			)
		}
	}

	return out
}

func (a *HTTPAgent) GenerateWithFixedSize(modelSpec types.ModelSpec, size int) (generator.DataGen, error) {
	out := make(generator.DataGen)

	if size <= 0 {
		return nil, errors.New("size must be greater than zero")
	}

	for key, value := range modelSpec {
		switch value.Dtype {
		case types.String:
			out[key] = generator.GenerateStringData(value.Constraint.StringConstraint.Values, size)
		case types.Int:
			out[key] = generator.GenerateIntegerData(
				value.Constraint.NumericConstraint.Mean,
				value.Constraint.NumericConstraint.StandardDeviation,
				size,
			)
		case types.Float:
			out[key] = generator.GenerateFloatData(
				value.Constraint.NumericConstraint.Mean,
				value.Constraint.NumericConstraint.StandardDeviation,
				size,
			)
		}
	}

	return out, nil
}

func (a *HTTPAgent) GenerateWithinRange(modelSpec types.ModelSpec, lowerBound int, upperBound int) (generator.DataGen, error) {
	out := make(generator.DataGen)

	if lowerBound <= 0 {
		return nil, errors.New("lower bound must be greater than zero")
	}

	if upperBound <= lowerBound {
		return nil, errors.New("higher bound must be greater than lower")
	}

	for key, value := range modelSpec {
		switch value.Dtype {
		case types.String:
			out[key] = generator.GenerateStringData(value.Constraint.StringConstraint.Values, upperBound)[lowerBound:upperBound]
		case types.Int:
			out[key] = generator.GenerateIntegerData(
				value.Constraint.NumericConstraint.Mean,
				value.Constraint.NumericConstraint.StandardDeviation,
				upperBound,
			)[lowerBound:upperBound]
		case types.Float:
			out[key] = generator.GenerateFloatData(
				value.Constraint.NumericConstraint.Mean,
				value.Constraint.NumericConstraint.StandardDeviation,
				upperBound,
			)[lowerBound:upperBound]
		}
	}

	return out, nil
}
