package pixela

import (
	"github.com/pkg/errors"
	"strings"
	"testing"
)

// TODO: introduction of test helper for table testing

// test for username validation
func TestValidator_usernameValidation(t *testing.T) {
	wantError := errors.New(validationErrorMessages["Username"])

	tests := []struct {
		name     string
		username string
		wantErr  error
	} {
		{"Normal case", "ho-2", nil},
		{"Empty", "", wantError},
		{"Upper case", "AA", wantError},
		{"Start with number", "0000", wantError},
		{"Start with hyphen", "-a", wantError},
		{"Too short", "a", wantError},
		{"Too long", strings.Repeat("a", 34), wantError},
	}

	type usernameValidator struct {
		Username string `validate:"username"`
	}

	validate := newValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vf := usernameValidator{Username: tt.username}
			gotErr := validate.Validate(vf)

			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but %#v", tt.wantErr, gotErr)
				}
			}
		})
	}
}

// test for token validation
func TestValidator_tokenValidation(t *testing.T) {
	wantError := errors.New(validationErrorMessages["Token"])

	tests := []struct {
		name    string
		token   string
		wantErr error
	} {
		{"Normal case", "testtoken", nil},
		{"Too short", strings.Repeat("a", 7), wantError},
		{"Too long", strings.Repeat("a", 129), wantError},
	}

	type tokenValidator struct {
		Token string `validate:"token"`
	}

	validate := newValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vf := tokenValidator{Token: tt.token}
			gotErr := validate.Validate(vf)

			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but %#v", tt.wantErr, gotErr)
				}
			}
		})
	}
}

// test for graphId validation
func TestValidator_graphIdValidation(t *testing.T) {
	wantError := errors.New(validationErrorMessages["GraphId"])

	tests := []struct {
		name    string
		graphId string
		wantErr error
	} {
		{"Normal case", "testtoken", nil},
		{"Empty", "", wantError},
		{"Upper case", "AA", wantError},
		{"Start with number", "0000", wantError},
		{"Start with hyphen", "-a", wantError},
		{"Too short", "a", wantError},
		{"Too long", strings.Repeat("a", 18), wantError},
	}

	type graphIdValidator struct {
		GraphId string `validate:"graphid"`
	}

	validate := newValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vf := graphIdValidator{GraphId: tt.graphId}
			gotErr := validate.Validate(vf)

			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but %#v", tt.wantErr, gotErr)
				}
			}
		})
	}
}

// test for date validation
func TestValidator_dateValidation(t *testing.T) {
	wantError := errors.New(validationErrorMessages["Date"])

	tests := []struct {
		name    string
		dateStr string
		wantErr error
	} {
		{"Normal case", "20190131", nil},
		{"Contains not digit", "a0000000", wantError},
		{"Lack of number of digits", "0000000", wantError},
		{"Over of number of digits", "000000000", wantError},
		{"Invalid month", "00009900", wantError},
		{"Invalid day", "00000099", wantError},
	}

	type dateValidator struct {
		Date string `validate:"date"`
	}

	validate := newValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vf := dateValidator{Date: tt.dateStr}
			gotErr := validate.Validate(vf)

			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but %#v", tt.wantErr, gotErr)
				}
			}
		})
	}
}

// test for quantity validation
func TestValidator_quantityValidation(t *testing.T) {
	wantError := errors.New(validationErrorMessages["Quantity"])

	tests := []struct {
		name        string
		quantityStr string
		wantErr     error
	} {
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

	type quantityValidator struct {
		Quantity string `validate:"quantity"`
	}

	validate := newValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vf := quantityValidator{Quantity: tt.quantityStr}
			gotErr := validate.Validate(vf)

			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but %#v", tt.wantErr, gotErr)
				}
			}
		})
	}
}

// test for optionalData validation
func TestValidator_optionalDataValidation(t *testing.T) {
	wantError := errors.New(validationErrorMessages["OptionalData"])

	tests := []struct {
		name        string
		optionalData string
		wantErr     error
	} {
		{"Normal case", `{"key":"value"}`, nil},
		{"Too long", strings.Repeat("a", 10241), wantError},
		{"Invalid JSON", `"key"`, wantError},
	}

	type optionalDataValidator struct {
		OptionalData string `validate:"optionaldata"`
	}

	validate := newValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vf := optionalDataValidator{OptionalData: tt.optionalData}
			gotErr := validate.Validate(vf)

			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but %#v", tt.wantErr, gotErr)
				}
			}
		})
	}
}
