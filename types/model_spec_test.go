package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseModelSpec(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    func() *ModelSpec
		wantErr bool
	}{
		{
			name: "successfully parsed model spec",
			jsonStr: `{
				"feature_name_1": {
					"dtype": "float",
					"constraint": {
						"numeric_constraint": {
							"std": 1.5,
							"mean": 2.0
						}
					}
				},
				"feature_name_2": {
					"dtype": "string",
					"constraint": {
						"string_constraint": {
							"values": ["a", "b"]
						}
					}
				}
			}`,
			want: func() *ModelSpec {
				want := make(ModelSpec)
				want["feature_name_1"] = Spec{
					Dtype: "float",
					Constraint: &Constraint{
						NumericConstraint: &NumericConstraint{
							StandardDeviation: 1.5,
							Mean:              2.0,
						},
						StringConstraint: nil,
					},
				}
				want["feature_name_2"] = Spec{
					Dtype: "string",
					Constraint: &Constraint{
						StringConstraint: &StringConstraint{Values: []string{"a", "b"}},
					},
				}
				return &want
			},
			wantErr: false,
		},
		{
			name: "successfully parsed model spec with numeric only constraints",
			jsonStr: `{
				"feature_name_1": {
					"dtype": "float",
					"constraint": {
						"numeric_constraint": {
							"std": 1.5,
							"mean": 2.0
						}
					}
				}
			}`,
			want: func() *ModelSpec {
				want := make(ModelSpec)
				want["feature_name_1"] = Spec{
					Dtype: "float",
					Constraint: &Constraint{
						NumericConstraint: &NumericConstraint{
							StandardDeviation: 1.5,
							Mean:              2.0,
						},
						StringConstraint: nil,
					},
				}
				return &want
			},
			wantErr: false,
		},
		{
			name: "successfully parsed model spec with string only constraints",
			jsonStr: `{
				"feature_name_1": {
					"dtype": "string",
					"constraint": {
						"string_constraint": {
							"values": ["a", "b"]
						}
					}
				}
			}`,
			want: func() *ModelSpec {
				want := make(ModelSpec)
				want["feature_name_1"] = Spec{
					Dtype: "string",
					Constraint: &Constraint{
						StringConstraint: &StringConstraint{Values: []string{"a", "b"}},
					},
				}
				return &want
			},
			wantErr: false,
		},
		{
			name:    "fails to parse incorrect model spec file",
			jsonStr: `incorrect json model spec file`,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseModelSpec([]byte(tt.jsonStr))
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want(), got)
			}
		})
	}
}
