package pixela

import (
	"github.com/pkg/errors"
	"strings"
	"testing"
)

type validateTestCase struct {
	name    string
	target  string
	wantErr error
}

type validateTestCases []validateTestCase

func validateTestHelper(t *testing.T, field string, tests validateTestCases) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var targetValidator interface{}

			switch field {
			case "Username":
				targetValidator = struct {
					target string `validate:"username"`
				}{
					tt.target,
				}
			case "Token":
				targetValidator = struct {
					target string `validate:"token"`
				}{
					tt.target,
				}
			case "GraphID":
				targetValidator = struct {
					target string `validate:"graphID"`
				}{
					tt.target,
				}
			case "Date":
				targetValidator = struct {
					target string `validate:"date"`
				}{
					tt.target,
				}
			case "Quantity":
				targetValidator = struct {
					target string `validate:"quantity"`
				}{
					tt.target,
				}
			case "OptionalData":
				targetValidator = struct {
					target string `validate:"optionaldata"`
				}{
					tt.target,
				}
			}

			validate := newValidator()
			gotErr := validate.Validate(targetValidator)

			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but %#v", tt.wantErr, gotErr)
				}
			}
		})
	}
}

// test for username validation
func TestValidator_usernameValidation(t *testing.T) {
	wantError := errors.New(validationErrorMessages["Username"])

	tests := validateTestCases{
		{"Normal case", "ho-2", nil},
		{"Empty", "", wantError},
		{"Upper case", "AA", wantError},
		{"Start with number", "0000", wantError},
		{"Start with hyphen", "-a", wantError},
		{"Too short", "a", wantError},
		{"Too long", strings.Repeat("a", 34), wantError},
	}

	validateTestHelper(t, "Username", tests)
}

// test for token validation
func TestValidator_tokenValidation(t *testing.T) {
	wantError := errors.New(validationErrorMessages["Token"])

	tests := validateTestCases{
		{"Normal case", "testtoken", nil},
		{"Too short", strings.Repeat("a", 7), wantError},
		{"Too long", strings.Repeat("a", 129), wantError},
	}

	validateTestHelper(t, "Token", tests)
}

// test for graphID validation
func TestValidator_graphIdValidation(t *testing.T) {
	wantError := errors.New(validationErrorMessages["GraphID"])

	tests := validateTestCases{
		{"Normal case", "testtoken", nil},
		{"Empty", "", wantError},
		{"Upper case", "AA", wantError},
		{"Start with number", "0000", wantError},
		{"Start with hyphen", "-a", wantError},
		{"Too short", "a", wantError},
		{"Too long", strings.Repeat("a", 18), wantError},
	}

	validateTestHelper(t, "GraphID", tests)
}

// test for date validation
func TestValidator_dateValidation(t *testing.T) {
	wantError := errors.New(validationErrorMessages["Date"])

	tests := validateTestCases{
		{"Normal case", "20190131", nil},
		{"Contains not digit", "a0000000", wantError},
		{"Lack of number of digits", "0000000", wantError},
		{"Over of number of digits", "000000000", wantError},
		{"Invalid month", "00009900", wantError},
		{"Invalid day", "00000099", wantError},
	}

	validateTestHelper(t, "Date", tests)
}

// test for quantity validation
func TestValidator_quantityValidation(t *testing.T) {
	wantError := errors.New(validationErrorMessages["Quantity"])

	tests := validateTestCases{
		{"Int (start with zero)", "0", nil},
		{"Int (start with none zero)", "1", nil},
		{"Int (with numbers of digits)", "11", nil},
		{"Float (succession with zero)", "0.0", nil},
		{"Float (start with zero)", "0.1", nil},
		{"Float (start with none zero)", "1.1", nil},
		{"Int (succession with zero)", "00", wantError},
		{"Include none digit char", "A", wantError},
		{"mix none digit char", "0A", wantError},
	}

	validateTestHelper(t, "Quantity", tests)
}

// test for optionalData validation
func TestValidator_optionalDataValidation(t *testing.T) {
	wantError := errors.New(validationErrorMessages["OptionalData"])

	tests := validateTestCases{
		{"Normal case", `{"key":"value"}`, nil},
		{"Normal case (empty string)", "", nil},
		{"Too long", strings.Repeat("a", 10241), wantError},
		{"Invalid JSON", `"key"`, wantError},
	}

	validateTestHelper(t, "OptionalData", tests)
}
