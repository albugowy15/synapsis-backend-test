package validator_test

import (
	"testing"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/validator"
)

func TestValidatePassword(t *testing.T) {
	testCase := []struct {
		name  string
		input string
		valid bool
	}{
		{
			name:  "has space",
			input: "feyf$ 7iY",
			valid: false,
		},
		{
			name:  "no uppercase",
			input: "fey7%jfee5",
			valid: false,
		},
		{
			name:  "no lowercase",
			input: "Z7%EEEH5",
			valid: false,
		},
		{
			name:  "no number",
			input: "Zy&jddedefeF",
			valid: false,
		},
		{
			name:  "no special char",
			input: "Zy43jddedefef",
			valid: false,
		},
		{
			name:  "valid",
			input: "V4lidP$ss",
			valid: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidatePassword(tc.input)
			if err != nil && tc.valid == true {
				t.Errorf("Expected valid password, got invalid")
			} else if err == nil && tc.valid == false {
				t.Errorf("Expected invalid password, got valid")
			}
		})
	}
}
