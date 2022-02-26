package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// EnvHandler keeps track of the errors internally and
// provides a method to obtain the errors as a combined error.
type EnvHandler struct {
	Errors []string
}

// NewEnvHandler creates a new builder with no errors.
func NewEnvHandler() *EnvHandler {
	return &EnvHandler{
		Errors: nil,
	}
}

// GetError returns the errors of this builder combined into a
// single error.
func (eh *EnvHandler) GetError() error {
	if len(eh.Errors) == 0 {
		return nil
	}
	return fmt.Errorf("configuration Error: %s", strings.Join(eh.Errors, ";"))
}

// GetUrl tries to extract a URL from the specified environment variable and
// takes an optional default value to return if the variable is not set.
// If the variable is set and cannot be parsed an error is noted in the builder.
// An error is also noted, if the variable is not set and no default value is
// given.
func (eh *EnvHandler) GetUrl(name string, defaultValue ...*url.URL) *url.URL {

	val := os.Getenv(name)
	if val == "" {
		if len(defaultValue) == 0 {
			eh.Errors = append(eh.Errors, fmt.Sprintf("missing value for %s", name))
			return nil
		} else {
			return defaultValue[0]
		}
	}
	u, err := url.Parse(val)
	if err != nil {
		eh.Errors = append(eh.Errors, fmt.Sprintf("cannot convert value %s to URL for var %s (%v)",
			val, name, err))
		return nil
	}
	return u
}

// GetInt tries to extract an integer from the specified environment variable and
// takes an optional default value to return if the variable is not set.
// If the variable is set and cannot be parsed to an integer an error is noted
// in the builder. An error is also noted, if the variable is not set and no default
// value is given.
func (eh *EnvHandler) GetInt(name string, defaultValue ...int) int {

	val := os.Getenv(name)
	if val == "" {
		if len(defaultValue) == 0 {
			eh.Errors = append(eh.Errors, fmt.Sprintf("missing value for %s", name))
			return 0
		} else {
			return defaultValue[0]
		}
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		eh.Errors = append(eh.Errors, fmt.Sprintf("cannot convert value %s to int for var %s (%v)",
			val, name, err))
		return 0
	}
	return i
}

// GetString tries to extract a string from the specified environment variable and
// takes an optional default value to return if the variable is not set.
// An error is noted, if the variable is not set and no default value is
// given.
func (eh *EnvHandler) GetString(name string, defaultValue ...string) string {

	val := os.Getenv(name)
	if val == "" {
		if len(defaultValue) == 0 {
			eh.Errors = append(eh.Errors, fmt.Sprintf("missing value for %s", name))
			return ""
		} else {
			return defaultValue[0]
		}
	}
	return val
}

// GetBool tries to extract a bool from the specified environment variable and
// takes an optional default value to return if the variable is not set.
// An error is noted, if the variable is not set and no default value is
// given.
func (eh *EnvHandler) GetBool(name string, defaultValue ...bool) bool {
	val := os.Getenv(name)
	if val == "" {
		if len(defaultValue) == 0 {
			eh.Errors = append(eh.Errors, fmt.Sprintf("missing value for %s", name))
			return false
		} else {
			return defaultValue[0]
		}
	}
	i, err := strconv.ParseBool(val)
	if err != nil {
		eh.Errors = append(eh.Errors, fmt.Sprintf("cannot convert value %s to bool for var %s (%v)",
			val, name, err))
		return false
	}
	return i
}

// GetOsHostname tries to extract the hostname provided by the
// operating system or records an error if not available.NewEnvHandler()
func (eh *EnvHandler) GetOsHostname() string {
	h, err := os.Hostname()
	if err != nil {
		eh.Errors = append(eh.Errors, fmt.Sprintf("cannot get hostname: %s", err.Error()))
		return ""
	}
	return h
}
