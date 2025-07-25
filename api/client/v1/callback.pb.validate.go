// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: client/v1/callback.proto

package clientv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on SendResultNotifyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SendResultNotifyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendResultNotifyRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendResultNotifyRequestMultiError, or nil if none found.
func (m *SendResultNotifyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SendResultNotifyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NotificationId

	if all {
		switch v := interface{}(m.GetRawRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SendResultNotifyRequestValidationError{
					field:  "RawRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SendResultNotifyRequestValidationError{
					field:  "RawRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRawRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SendResultNotifyRequestValidationError{
				field:  "RawRequest",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SendResultNotifyRequestValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SendResultNotifyRequestValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SendResultNotifyRequestValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SendResultNotifyRequestMultiError(errors)
	}

	return nil
}

// SendResultNotifyRequestMultiError is an error wrapping multiple validation
// errors returned by SendResultNotifyRequest.ValidateAll() if the designated
// constraints aren't met.
type SendResultNotifyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendResultNotifyRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendResultNotifyRequestMultiError) AllErrors() []error { return m }

// SendResultNotifyRequestValidationError is the validation error returned by
// SendResultNotifyRequest.Validate if the designated constraints aren't met.
type SendResultNotifyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendResultNotifyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendResultNotifyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendResultNotifyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendResultNotifyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendResultNotifyRequestValidationError) ErrorName() string {
	return "SendResultNotifyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SendResultNotifyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendResultNotifyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendResultNotifyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendResultNotifyRequestValidationError{}

// Validate checks the field values on SendResultNotifyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SendResultNotifyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendResultNotifyResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendResultNotifyResponseMultiError, or nil if none found.
func (m *SendResultNotifyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SendResultNotifyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return SendResultNotifyResponseMultiError(errors)
	}

	return nil
}

// SendResultNotifyResponseMultiError is an error wrapping multiple validation
// errors returned by SendResultNotifyResponse.ValidateAll() if the designated
// constraints aren't met.
type SendResultNotifyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendResultNotifyResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendResultNotifyResponseMultiError) AllErrors() []error { return m }

// SendResultNotifyResponseValidationError is the validation error returned by
// SendResultNotifyResponse.Validate if the designated constraints aren't met.
type SendResultNotifyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendResultNotifyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendResultNotifyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendResultNotifyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendResultNotifyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendResultNotifyResponseValidationError) ErrorName() string {
	return "SendResultNotifyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SendResultNotifyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendResultNotifyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendResultNotifyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendResultNotifyResponseValidationError{}
