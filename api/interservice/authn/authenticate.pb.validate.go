// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: interservice/authn/authenticate.proto

package authn

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _authenticate_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AuthenticateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AuthenticateRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// AuthenticateRequestValidationError is the validation error returned by
// AuthenticateRequest.Validate if the designated constraints aren't met.
type AuthenticateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthenticateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthenticateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthenticateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthenticateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthenticateRequestValidationError) ErrorName() string {
	return "AuthenticateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AuthenticateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthenticateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthenticateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthenticateRequestValidationError{}

// Validate checks the field values on AuthenticateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AuthenticateResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Subject

	return nil
}

// AuthenticateResponseValidationError is the validation error returned by
// AuthenticateResponse.Validate if the designated constraints aren't met.
type AuthenticateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthenticateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthenticateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthenticateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthenticateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthenticateResponseValidationError) ErrorName() string {
	return "AuthenticateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthenticateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthenticateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthenticateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthenticateResponseValidationError{}
