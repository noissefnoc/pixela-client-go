package pixela

import (
	"encoding/json"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
	"regexp"
	"strings"
	"time"
)

type validateField struct {
	Username            string `validate:"omitempty,username"`
	Token               string `validate:"omitempty,token"`
	AgreeTermsOfService string `validate:"omitempty,oneof=yes no"`
	NotMinor            string `validate:"omitempty,oneof=yes no"`
	NewToken            string `validate:"omitempty,token"`
	GraphId             string `validate:"omitempty,graphid"`
	UnitType            string `validate:"omitempty,oneof=int float"`
	Color               string `validate:"omitempty,oneof=shibafu momiji sora ichou ajisai kuro"`
	Date                string `validate:"omitempty,date"`
	Quantity            string `validate:"omitempty,quantity"`
	WebhookType         string `validate:"omitempty,oneof=increment decrement"`
	OptionalData        string `validate:"omitempty,optionaldata"`
	SelfSufficient      string `validate:"omitempty,oneof=none increment decrement"`
}

type Validator struct {
	validator *validator.Validate
}

func newValidator() Validator {
	validate := validator.New()

	validate.RegisterValidation("username", usernameValidation)
	validate.RegisterValidation("token", tokenValidation)
	validate.RegisterValidation("graphid", graphIdValidator)
	validate.RegisterValidation("date", dateValidator)
	validate.RegisterValidation("quantity", quantityValidator)
	validate.RegisterValidation("optionaldata", optionalDataValidator)

	return Validator{validator: validate}
}

var validationErrorMessages = map[string]string{
	"Username": "`username` allows lowercase alphabet, number and hyphen (NOTE: first letter only allows alphabet.) and 1 to 32 length.",
	"Token": "`token` allows 8 to 128 length.",
	"AgreeTermsOfService": "`agreeTermsOfService` allows `yes` or `no`.",
	"NotMinor": "`notMinor` allows `yes` or `no`.",
	"NewToken": "`newToken` allows 8 to 128 length.",
	"GraphId": "`graphId` allows lowercase alphabet, number and hyphen (NOTE: first letter only allows alphabet.) and 1 to 16 length.",
	"UnitType": "`unit` allows `int` or `float`.",
	"Color": "`color` allows `shibafu`, `momiji`, `sora`, `ichou`, `ajisai` or `kuro`.",
	"Date": "`date` format is `yyyyMMdd`.",
	"Quantity": "`quantity` allows value of int or float.",
	"WebhookType": "`type` allows `increment` or `decrement`.",
	"OptionalData": "`optionalData` is under 10k JSON string.",
	"SelfSufficient": "`selfSufficient` allows `increment` or `decrement`.",
}

func (pv *Validator) Validate(i interface{}) error {
	err := pv.validator.Struct(i)

	var errorMessages []string

	if err != nil {
		for _, err := range err.(validator.ValidationErrors){
			errorMessages = append(errorMessages, validationErrorMessages[err.Field()])
		}

		return errors.New(strings.Join(errorMessages, " and "))
	}

	return nil
}

// username validator
func usernameValidation(fl validator.FieldLevel) bool {
	tf, err := regexp.Match(`^[a-z][a-z0-9-]{1,32}$`, []byte(fl.Field().String()))

	if tf && err == nil {
		return true
	}

	return false
}

// token validator
func tokenValidation(fl validator.FieldLevel) bool {
	tf, err := regexp.Match(`^[ -~]{8,128}$`, []byte(fl.Field().String()))

	if tf && err == nil {
		return true
	}

	return false
}

// graphId validator
func graphIdValidator(fl validator.FieldLevel) bool {
	tf, err := regexp.Match(`^[a-z][a-z0-9-]{1,16}$`, []byte(fl.Field().String()))

	if tf && err == nil {
		return true
	}

	return false
}

// date validator
func dateValidator(fl validator.FieldLevel) bool {
	_, err := time.Parse("20060102", fl.Field().String())

	if err != nil {
		return false
	}

	return true
}

// quantity validator
func quantityValidator(fl validator.FieldLevel) bool {
	tf, err := regexp.Match(`^(0|[1-9][0-9]*)(\.[0-9])*$`, []byte(fl.Field().String()))

	if tf && err == nil {
		return true
	}

	return false
}

// optionalData validator
func optionalDataValidator(fl validator.FieldLevel) bool {
	optionalData := fl.Field().String()

	// empty string
	if len(optionalData) == 0 {
		return true
	}

	if len(optionalData) > 10240 || !json.Valid([]byte(optionalData)) {
		return false
	}

	return true
}
