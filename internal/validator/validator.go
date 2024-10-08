package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRX = regexp.MustCompile(
	`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`,
)

// Define Validator type which contains a map of validation errors for form fields
type Validator struct {
	NonFieldErrors []string
	FieldErrors    map[string]string
}

// Valid() return true if the FieldErrors map doesn't contain any entries
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

// AddFieldError() adds an error message to the FieldErrors map (so long as no
// entry already exists for the given key)
func (v *Validator) AddFieldError(key, message string) {
	// Need to initialize the map first, if it isn't already initialized
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// AddNonFieldError() helper for adding error messages to the new NonFieldErrors slice
func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

// CheckField() adds an error message to the FieldErrors map only if a validation check is not 'ok'
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// NotBlank() returns true if a value is not an empty string
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars() returns true if a value contains no more than n characters
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// PermittedInt() returns true if a value is in a list of permitted integers
func PermittedInt(value int, permittedValue ...int) bool {
	for i := range permittedValue {
		if value == permittedValue[i] {
			return true
		}
	}
	return false
}

func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func Mathes(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func PermittedValue[T comparable](value T, permittedValue ...T) bool {
	for i := range permittedValue {
		if value == permittedValue[i] {
			return true
		}
	}
	return false
}
