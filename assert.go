package argo

import (
	"fmt"
	"strings"
)

func illegalArgumentIf(condition bool, msg string, args ...interface{}) error {
	if condition {
		return &IllegalArgumentError{fmt.Sprintf(msg, args...)}
	}
	return nil
}

func illegalStateIf(condition bool, msg string, args ...interface{}) error {
	if condition {
		return &IllegalStateError{fmt.Sprintf(msg, args...)}
	}
	return nil
}

// State will check if supplied state condition is true.
func State(state bool, msg string, args ...interface{}) error {
	return illegalStateIf(!state, msg, args...)
}

// Is will check that supplied condition is true and return an error if not.
func Is(value bool, msg string, args ...interface{}) error {
	return illegalArgumentIf(!value, msg, args...)
}

// Not will check that supplied condition is false and return an error if not.
func Not(value bool, msg string, args ...interface{}) error {
	return illegalArgumentIf(value, msg, args...)
}

// NotNil will check if supplied value is not nil.
func NotNil(value interface{}, argName string) error {
	return illegalArgumentIf(value == nil, "%s is required", argName)
}

// NotBlank will check if supplied string value is not blank.
func NotBlank(value string, argName string) error {
	value = strings.TrimSpace(value)
	return illegalArgumentIf(len(value) == 0, "%s is expected to be not blank", argName)
}

// Length will check if supplied string has exactly supplied length.
func Length(value string, exactLength int, argName string) error {
	length := len(value)
	return illegalArgumentIf(length == exactLength, "%s must be long exactly %d", argName, exactLength)
}

// MinLength will check if supplied string has provided minimum length.
func MinLength(value string, minLength int, argName string) error {
	length := len(value)
	return illegalArgumentIf(length < minLength, "%s must be long at least %d", argName, minLength)
}

// MaxLength will check if supplied string has provided maximum length.
func MaxLength(value string, maxLength int, argName string) error {
	length := len(value)
	return illegalArgumentIf(length > maxLength, "%s must be long at most %d", argName, maxLength)
}

// LengthBetween will check if supplied string has length between min and max.
func LengthBetween(value string, minLength, maxLength int, argName string) error {
	length := len(value)
	return illegalArgumentIf(length < minLength || length > maxLength, "%s length must be between %d and %d",
		argName, minLength, maxLength)
}

// IsTrue check that supplied value is true.
func IsTrue(value bool, argName string) error {
	return illegalArgumentIf(value != true, "%s is expected to be true", argName)
}

// IsFalse check that supplied value is false.
func IsFalse(value bool, argName string) error {
	return illegalArgumentIf(value != false, "%s is expected to be false", argName)
}
