package pixela

import "testing"

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid username", "test-user-1", false},
		{"valid single char + alphanumeric", "a1", false},
		{"starts with number", "1test", true},
		{"starts with hyphen", "-test", true},
		{"uppercase", "Test", true},
		{"too long", "abcdefghijklmnopqrstuvwxyz1234567", true},
		{"empty", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateUsername(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateUsername(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateToken(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid token 8 chars", "12345678", false},
		{"valid token with symbols", "abc-def_ghi", false},
		{"too short", "1234567", true},
		{"empty", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateToken(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateToken(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateGraphID(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid graph ID", "test-graph-1", false},
		{"starts with number", "1graph", true},
		{"starts with hyphen", "-graph", true},
		{"uppercase", "Graph", true},
		{"too long", "abcdefghijklmnopq", true},
		{"empty", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateGraphID(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateGraphID(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateDate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid date", "20261102", false},
		{"invalid format with hyphen", "2026-11-02", true},
		{"invalid month", "20261301", true},
		{"invalid day", "20260230", true},
		{"empty", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateDate(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateQuantity(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid integer zero", "0", false},
		{"valid integer non-zero", "100", false},
		{"valid float", "10.5", false},
		{"valid float with zero prefix", "0.5", false},
		{"invalid string", "abc", true},
		{"empty", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateQuantity(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateQuantity(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateColor(t *testing.T) {
	validColors := []string{"shibafu", "momiji", "sora", "ichou", "ajisai", "kuro"}
	for _, c := range validColors {
		if err := validateColor(c); err != nil {
			t.Errorf("validateColor(%q) unexpected error: %v", c, err)
		}
	}

	if err := validateColor("invalid"); err == nil {
		t.Errorf("validateColor('invalid') expected error, got nil")
	}
}

func TestValidateUnitType(t *testing.T) {
	for _, u := range []string{"int", "float"} {
		if err := validateUnitType(u); err != nil {
			t.Errorf("validateUnitType(%q) unexpected error: %v", u, err)
		}
	}

	if err := validateUnitType("string"); err == nil {
		t.Errorf("validateUnitType('string') expected error, got nil")
	}
}

func TestValidateSelfSufficient(t *testing.T) {
	for _, ss := range []string{"", "none", "increment", "decrement"} {
		if err := validateSelfSufficient(ss); err != nil {
			t.Errorf("validateSelfSufficient(%q) unexpected error: %v", ss, err)
		}
	}

	if err := validateSelfSufficient("invalid"); err == nil {
		t.Errorf("validateSelfSufficient('invalid') expected error, got nil")
	}
}

func TestValidateWebhookType(t *testing.T) {
	for _, wt := range []string{"increment", "decrement"} {
		if err := validateWebhookType(wt); err != nil {
			t.Errorf("validateWebhookType(%q) unexpected error: %v", wt, err)
		}
	}

	if err := validateWebhookType("invalid"); err == nil {
		t.Errorf("validateWebhookType('invalid') expected error, got nil")
	}
}

func TestValidateOptionalData(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"empty", "", false},
		{"valid JSON", `{"key":"value"}`, false},
		{"invalid JSON", `not-json`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateOptionalData(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateOptionalData(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}
