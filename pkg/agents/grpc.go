package agents

import (
	"errors"
	"math/rand"

	"github.com/gagansingh894/InferFlux/types"
)

type GRPCAgent struct{}

func NewGRPCAgent() (*GRPCAgent, error) {
	return &GRPCAgent{}, nil
}

func (a *GRPCAgent) GenerateWithRandomSize(modelSpec types.ModelSpec) DataGen {
	out := make(DataGen)
	size := rand.Intn(1500) + 1

	for key, value := range modelSpec {
		switch value.Dtype {
		case types.String:
			out[key] = generateStringData(value.Constraint.StringConstraint.Values, size)
		case types.Int:
			out[key] = generateIntegerData(
				value.Constraint.NumericConstraint.Mean,
				value.Constraint.NumericConstraint.StandardDeviation,
				size,
			)
		case types.Float:
			out[key] = generateFloatData(
				value.Constraint.NumericConstraint.Mean,
				value.Constraint.NumericConstraint.StandardDeviation,
				size,
			)
		}
	}

	return out
}

func (a *GRPCAgent) GenerateWithFixedSize(modelSpec types.ModelSpec, size int) (DataGen, error) {
	out := make(DataGen)

	if size <= 0 {
		return nil, errors.New("size must be greater than zero")
	}

	for key, value := range modelSpec {
		switch value.Dtype {
		case types.String:
			out[key] = generateStringData(value.Constraint.StringConstraint.Values, size)
		case types.Int:
			out[key] = generateIntegerData(
				value.Constraint.NumericConstraint.Mean,
				value.Constraint.NumericConstraint.StandardDeviation,
				size,
			)
		case types.Float:
			out[key] = generateFloatData(
				value.Constraint.NumericConstraint.Mean,
				value.Constraint.NumericConstraint.StandardDeviation,
				size,
			)
		}
	}

	return out, nil
}

func (a *GRPCAgent) GenerateWithinRange(modelSpec types.ModelSpec, lowerBound int, upperBound int) (DataGen, error) {
	out := make(DataGen)

	if lowerBound <= 0 {
		return nil, errors.New("lower bound must be greater than zero")
	}

	if upperBound <= lowerBound {
		return nil, errors.New("higher bound must be greater than lower")
	}

	for key, value := range modelSpec {
		switch value.Dtype {
		case types.String:
			out[key] = generateStringData(value.Constraint.StringConstraint.Values, upperBound)[lowerBound:upperBound]
		case types.Int:
			out[key] = generateIntegerData(
				value.Constraint.NumericConstraint.Mean,
				value.Constraint.NumericConstraint.StandardDeviation,
				upperBound,
			)[lowerBound:upperBound]
		case types.Float:
			out[key] = generateFloatData(
				value.Constraint.NumericConstraint.Mean,
				value.Constraint.NumericConstraint.StandardDeviation,
				upperBound,
			)[lowerBound:upperBound]
		}
	}

	return out, nil
}
