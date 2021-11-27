// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service.proto

package note_v1

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

// Validate checks the field values on CreateNoteV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateNoteV1Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNoteV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateNoteV1RequestMultiError, or nil if none found.
func (m *CreateNoteV1Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNoteV1Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for ClassroomId

	// no validation rules for DocumentId

	if len(errors) > 0 {
		return CreateNoteV1RequestMultiError(errors)
	}
	return nil
}

// CreateNoteV1RequestMultiError is an error wrapping multiple validation
// errors returned by CreateNoteV1Request.ValidateAll() if the designated
// constraints aren't met.
type CreateNoteV1RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNoteV1RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNoteV1RequestMultiError) AllErrors() []error { return m }

// CreateNoteV1RequestValidationError is the validation error returned by
// CreateNoteV1Request.Validate if the designated constraints aren't met.
type CreateNoteV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNoteV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNoteV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNoteV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNoteV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNoteV1RequestValidationError) ErrorName() string {
	return "CreateNoteV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNoteV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNoteV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNoteV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNoteV1RequestValidationError{}

// Validate checks the field values on CreateNoteV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateNoteV1Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNoteV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateNoteV1ResponseMultiError, or nil if none found.
func (m *CreateNoteV1Response) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNoteV1Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NoteId

	if len(errors) > 0 {
		return CreateNoteV1ResponseMultiError(errors)
	}
	return nil
}

// CreateNoteV1ResponseMultiError is an error wrapping multiple validation
// errors returned by CreateNoteV1Response.ValidateAll() if the designated
// constraints aren't met.
type CreateNoteV1ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNoteV1ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNoteV1ResponseMultiError) AllErrors() []error { return m }

// CreateNoteV1ResponseValidationError is the validation error returned by
// CreateNoteV1Response.Validate if the designated constraints aren't met.
type CreateNoteV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNoteV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNoteV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNoteV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNoteV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNoteV1ResponseValidationError) ErrorName() string {
	return "CreateNoteV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNoteV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNoteV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNoteV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNoteV1ResponseValidationError{}
