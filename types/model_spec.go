package types

import (
	"encoding/json"
	"fmt"
)

// ModelSpec represents the specification of a model , including the
// feature's name and it's detailed specification
type ModelSpec map[string]Spec

// Spec defines the datatype and optional constraints for a model feature.
// It includes fields for specifying numeric boundaries or valid string values
// based on the feature's data type.
type Spec struct {
	// Dtype specifies the data type of the feature (e.g., "float", "int" or "string").
	Dtype string `json:"dtype"`
	// Constraint optionally specifies additional limits and conditions for
	// the feature's values. If unset, it will be omitted from JSON output.
	Constraint *Constraint `json:"constraint,omitempty"`
}

// Constraint defines optional constraints that may apply to a feature
// with a specific data type. It includes numeric boundaries, intervals,
// and valid values for string types. This struct can be used to enforce
// rules on the input data during validation or preprocessing.
type Constraint struct {
	// NumericConstraint specifies constraints for numeric data types,
	// such as mean and standard deviation.
	NumericConstraint *NumericConstraint `json:"numeric_constraint,omitempty"`

	// StringConstraint specifies constraints for string data types,
	// such as a set of acceptable values.
	StringConstraint *StringConstraint `json:"string_constraint,omitempty"`
}

// NumericConstraint defines constraints for numeric feature data.
// These constraints include statistical properties such as the mean
// and standard deviation, which can be used for validation or normalization.
type NumericConstraint struct {
	// StandardDeviation represents the standard deviation of the feature data.
	// This value is used for validation or scaling purposes.
	StandardDeviation float64 `json:"std"`

	// Mean represents the average (mean) value of the feature data.
	// This value is used for validation or scaling purposes.
	Mean float64 `json:"mean"`
}

// StringConstraint defines constraints for string feature data.
// These constraints include a list of acceptable values to ensure
// input data adheres to a predefined set of valid strings.
type StringConstraint struct {
	// Values is a list of acceptable string values for the feature.
	// This field is relevant only when the data type of the feature is "string".
	// It is ignored for numeric types.
	Values []string `json:"values"`
}

// ParseModelSpec parses a JSON-encoded ModelSpec from the provided byte slice.
// It attempts to unmarshal the data into a ModelSpec struct, returning an
// error if the data is invalid or improperly formatted.
//
// Parameters:
//   - data: A byte slice containing the JSON representation of a ModelSpec.
//
// Returns:
//   - (*ModelSpec, error): A pointer to a ModelSpec struct if parsing succeeds,
//     or an error describing the parsing failure.
func ParseModelSpec(data []byte) (*ModelSpec, error) {
	var modelSpec ModelSpec

	// Attempt to unmarshal JSON data into the modelSpec struct
	err := json.Unmarshal(data, &modelSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to parse model spec: %w", err)
	}

	// Validate model spec
	err = modelSpec.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate model spec: %w", err)
	}

	return &modelSpec, nil
}
