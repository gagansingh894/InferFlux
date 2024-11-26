package types

import (
	"fmt"
	"slices"
)

const (
	// Float represents the "float" data type.
	Float string = "float"

	// String represents the "string" data type.
	String string = "string"

	// Int represents the "int" data type.
	Int string = "int"
)

// validator defines a function type that validates a ModelSpec.
// A validator function returns an error if validation fails.
type validator func(m ModelSpec) error

// Validate validates the ModelSpec using a series of validators.
// It runs each validator function sequentially and returns an error if any validation fails.
func (m ModelSpec) Validate() error {
	validators := []validator{
		validateDType,       // Validates that the data type is valid.
		validateConstraints, // Validates that constraints are appropriately defined.
	}

	for _, v := range validators {
		if err := v(m); err != nil {
			return fmt.Errorf("validation error: %v", err)
		}
	}

	return nil
}

// validateDType checks if the data types of all features in the ModelSpec are valid.
// Returns an error if any invalid data type is found.
func validateDType(m ModelSpec) error {
	for _, value := range m {
		// Check if the dtype is in the list of valid data types.
		found := slices.Contains(validDTypes(), value.Dtype)
		if !found {
			return fmt.Errorf("dtype %s is not valid", value.Dtype)
		}
	}

	return nil
}

// validateConstraints checks the constraints of each feature in the ModelSpec.
// Ensures that:
// - A feature does not have both numeric and string constraints simultaneously.
// - A feature has at least one constraint.
// - A string feature has string constraints.
// - A numeric feature has numeric constraints.
// Returns an error if any of these rules are violated.
func validateConstraints(m ModelSpec) error {
	for key, value := range m {
		if value.Constraint.NumericConstraint != nil && value.Constraint.StringConstraint != nil {
			return fmt.Errorf("feature: %s has both numeric constraint and string constraint", key)
		}

		if value.Constraint.NumericConstraint == nil && value.Constraint.StringConstraint == nil {
			return fmt.Errorf("feature: %s has no constraints", key)
		}

		if value.Dtype == String && value.Constraint.StringConstraint == nil {
			return fmt.Errorf("feature: %s is string but does not have any string constraint", key)
		}

		if value.Dtype != String && value.Constraint.NumericConstraint == nil {
			return fmt.Errorf("feature: %s is numeric but does not have any numeric constraint", key)
		}
	}

	return nil
}

// validDTypes returns a list of valid data types.
// These are the allowed types for features in the ModelSpec.
func validDTypes() []string {
	return []string{Float, String, Int}
}
