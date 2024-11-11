package types

import (
	"encoding/json"
	"fmt"
)

// ModelSpec represents the specification of a model , including the
// feature's name and it's detailed specification
type ModelSpec struct {
	// FeatureName is the name of the feature in the model.
	FeatureName string `json:"feature_name"`
	// Spec defines the datatype and optional constraints associated with
	// the feature. This struct includes the dtype (e.g., "float", "string")
	// and an optional Constraint that further limits valid values.
	Spec Spec `json:"spec"`
}

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

// Constraint defines optional constraints that may apply to a feature with a
// specific data type. It includes numeric boundaries, intervals, and valid
// values for string types.
type Constraint struct {
	// Min is the minimum allowable value for numeric types.
	Min float64 `json:"min"`

	// Max is the maximum allowable value for numeric types.
	Max float64 `json:"max"`

	// Interval is the step size for numeric types, indicating valid increments.
	Interval float64 `json:"interval"`

	// Values is a list of acceptable values, applicable when Dtype is "string".
	// This field is ignored for numeric types.
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
	return &modelSpec, nil
}
