package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelSpecValidate(t *testing.T) {
	tests := []struct {
		name      string
		modelSpec func() ModelSpec
		wantErr   bool
	}{
		{
			name: "valid model spec",
			modelSpec: func() ModelSpec {
				want := make(ModelSpec)
				want["feature_name_1"] = Spec{
					Dtype: "float",
					Constraint: &Constraint{
						NumericConstraint: &NumericConstraint{
							StandardDeviation: 1.5,
							Mean:              2.0,
						},
					},
				}
				want["feature_name_2"] = Spec{
					Dtype: "string",
					Constraint: &Constraint{
						StringConstraint: &StringConstraint{Values: []string{"a", "b"}},
					},
				}
				return want
			},
		},
		{
			name: "invalid dtype",
			modelSpec: func() ModelSpec {
				want := make(ModelSpec)
				want["feature_name_1"] = Spec{
					Dtype: "invalid",
				}
				return want
			},
			wantErr: true,
		},
		{
			name: "invalid constraint - both numerical and string constraints",
			modelSpec: func() ModelSpec {
				want := make(ModelSpec)
				want["feature_name_1"] = Spec{
					Dtype: "float",
					Constraint: &Constraint{
						NumericConstraint: &NumericConstraint{
							StandardDeviation: 1.5,
							Mean:              2.0,
						},
						StringConstraint: &StringConstraint{Values: []string{"a", "b"}},
					},
				}
				return want
			},
			wantErr: true,
		},
		{
			name: "invalid constraint - no constraint",
			modelSpec: func() ModelSpec {
				want := make(ModelSpec)
				want["feature_name_1"] = Spec{
					Dtype: "float",
					Constraint: &Constraint{
						NumericConstraint: nil,
						StringConstraint:  nil,
					},
				}
				return want
			},
			wantErr: true,
		},
		{
			name: "invalid constraint- string dtype but no corresponding constraint",
			modelSpec: func() ModelSpec {
				want := make(ModelSpec)
				want["feature_name_1"] = Spec{
					Dtype: "string",
					Constraint: &Constraint{
						NumericConstraint: &NumericConstraint{
							StandardDeviation: 1.5,
							Mean:              2.0,
						},
						StringConstraint: nil,
					},
				}
				return want
			},
			wantErr: true,
		},
		{
			name: "invalid constraint- numeric dtype but no corresponding constraint",
			modelSpec: func() ModelSpec {
				want := make(ModelSpec)
				want["feature_name_1"] = Spec{
					Dtype: "float",
					Constraint: &Constraint{
						NumericConstraint: nil,
						StringConstraint:  &StringConstraint{Values: []string{"a", "b"}},
					},
				}
				return want
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.modelSpec().Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
