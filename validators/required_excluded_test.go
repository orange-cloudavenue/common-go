package validators

import "testing"

func TestValidator_RequireIfNull(t *testing.T) {
	type data struct {
		Field1 string `validate:"required_if_null=Field2"`
		Field2 string
	}

	tests := []struct {
		name        string
		input       data
		expectedErr bool
	}{
		{name: "Field2 is null, Field1 is set", input: data{Field1: "value", Field2: ""}, expectedErr: false},
		{name: "Field2 is null, Field1 is null", input: data{Field1: "", Field2: ""}, expectedErr: true},
		{name: "Field2 is set, Field1 is set", input: data{Field1: "value", Field2: "not null"}, expectedErr: false},
		{name: "Field2 is set, Field1 is null", input: data{Field1: "", Field2: "not null"}, expectedErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := New().Struct(&tt.input)
			if (err == nil) == tt.expectedErr {
				t.Errorf("expected validation result %v, got %v", tt.expectedErr, err == nil)
			}
		})
	}
}

func TestValidator_ExcludeIfNull(t *testing.T) {
	type data struct {
		Field1 string `validate:"excluded_if_null=Field2"`
		Field2 string
	}

	tests := []struct {
		name        string
		input       data
		expectedErr bool
	}{
		{name: "Field2 is null, Field1 is set", input: data{Field1: "value", Field2: ""}, expectedErr: true},
		{name: "Field2 is null, Field1 is null", input: data{Field1: "", Field2: ""}, expectedErr: false},
		{name: "Field2 is set, Field1 is set", input: data{Field1: "value", Field2: "not null"}, expectedErr: false},
		{name: "Field2 is set, Field1 is null", input: data{Field1: "", Field2: "not null"}, expectedErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := New().Struct(&tt.input)
			if (err == nil) == tt.expectedErr {
				t.Errorf("expected validation result %v, got %v", tt.expectedErr, err == nil)
			}
		})
	}
}
