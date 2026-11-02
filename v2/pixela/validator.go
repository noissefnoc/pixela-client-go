package pixela

import (
	"encoding/json"
	"errors"
	"regexp"
	"time"
)

var (
	usernameRegex = regexp.MustCompile(`^[a-z][a-z0-9-]{0,31}$`)
	tokenRegex    = regexp.MustCompile(`^[ -~]{8,128}$`)
	graphIDRegex  = regexp.MustCompile(`^[a-z][a-z0-9-]{0,15}$`)
	quantityRegex = regexp.MustCompile(`^(0|[1-9][0-9]*)(\.[0-9])*$`)
)

func validateUsername(username string) error {
	if !usernameRegex.MatchString(username) {
		return errors.New("`username` allows lowercase alphabet, number and hyphen (NOTE: first letter only allows alphabet.) and 1 to 32 length.")
	}
	return nil
}

func validateToken(token string) error {
	if !tokenRegex.MatchString(token) {
		return errors.New("`token` allows 8 to 128 length.")
	}
	return nil
}

func validateGraphID(graphID string) error {
	if !graphIDRegex.MatchString(graphID) {
		return errors.New("`graphID` allows lowercase alphabet, number and hyphen (NOTE: first letter only allows alphabet.) and 1 to 16 length.")
	}
	return nil
}

func validateDate(date string) error {
	if _, err := time.Parse("20060102", date); err != nil {
		return errors.New("`date` format is `yyyyMMdd`.")
	}
	return nil
}

func validateQuantity(quantity string) error {
	if !quantityRegex.MatchString(quantity) {
		return errors.New("`quantity` allows value of int or float.")
	}
	return nil
}

func validateColor(color string) error {
	switch color {
	case "shibafu", "momiji", "sora", "ichou", "ajisai", "kuro":
		return nil
	default:
		return errors.New("`color` allows `shibafu`, `momiji`, `sora`, `ichou`, `ajisai` or `kuro`.")
	}
}

func validateUnitType(unitType string) error {
	switch unitType {
	case "int", "float":
		return nil
	default:
		return errors.New("`unit` allows `int` or `float`.")
	}
}

func validateSelfSufficient(ss string) error {
	if ss == "" {
		return nil
	}
	switch ss {
	case "none", "increment", "decrement":
		return nil
	default:
		return errors.New("`selfSufficient` allows `increment` or `decrement`.")
	}
}

func validateWebhookType(typ string) error {
	switch typ {
	case "increment", "decrement":
		return nil
	default:
		return errors.New("`type` allows `increment` or `decrement`.")
	}
}

func validateOptionalData(data string) error {
	if data == "" {
		return nil
	}
	if len(data) > 10240 || !json.Valid([]byte(data)) {
		return errors.New("`optionalData` is under 10k JSON string.")
	}
	return nil
}
