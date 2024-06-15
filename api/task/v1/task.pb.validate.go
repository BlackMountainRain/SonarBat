// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/task/v1/task.proto

package v1

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

// Validate checks the field values on CreateTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTaskRequestMultiError, or nil if none found.
func (m *CreateTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Name

	// no validation rules for Comment

	if len(errors) > 0 {
		return CreateTaskRequestMultiError(errors)
	}

	return nil
}

// CreateTaskRequestMultiError is an error wrapping multiple validation errors
// returned by CreateTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTaskRequestMultiError) AllErrors() []error { return m }

// CreateTaskRequestValidationError is the validation error returned by
// CreateTaskRequest.Validate if the designated constraints aren't met.
type CreateTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTaskRequestValidationError) ErrorName() string {
	return "CreateTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTaskRequestValidationError{}

// Validate checks the field values on CreateTaskReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateTaskReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTaskReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTaskReplyMultiError, or nil if none found.
func (m *CreateTaskReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTaskReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateTaskReplyMultiError(errors)
	}

	return nil
}

// CreateTaskReplyMultiError is an error wrapping multiple validation errors
// returned by CreateTaskReply.ValidateAll() if the designated constraints
// aren't met.
type CreateTaskReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTaskReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTaskReplyMultiError) AllErrors() []error { return m }

// CreateTaskReplyValidationError is the validation error returned by
// CreateTaskReply.Validate if the designated constraints aren't met.
type CreateTaskReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTaskReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTaskReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTaskReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTaskReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTaskReplyValidationError) ErrorName() string { return "CreateTaskReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateTaskReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTaskReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTaskReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTaskReplyValidationError{}

// Validate checks the field values on UpdateTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTaskRequestMultiError, or nil if none found.
func (m *UpdateTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.Name != nil {
		// no validation rules for Name
	}

	if m.Comment != nil {
		// no validation rules for Comment
	}

	if len(errors) > 0 {
		return UpdateTaskRequestMultiError(errors)
	}

	return nil
}

// UpdateTaskRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTaskRequestMultiError) AllErrors() []error { return m }

// UpdateTaskRequestValidationError is the validation error returned by
// UpdateTaskRequest.Validate if the designated constraints aren't met.
type UpdateTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTaskRequestValidationError) ErrorName() string {
	return "UpdateTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTaskRequestValidationError{}

// Validate checks the field values on UpdateTaskReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateTaskReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTaskReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTaskReplyMultiError, or nil if none found.
func (m *UpdateTaskReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTaskReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RowsAffected

	if len(errors) > 0 {
		return UpdateTaskReplyMultiError(errors)
	}

	return nil
}

// UpdateTaskReplyMultiError is an error wrapping multiple validation errors
// returned by UpdateTaskReply.ValidateAll() if the designated constraints
// aren't met.
type UpdateTaskReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTaskReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTaskReplyMultiError) AllErrors() []error { return m }

// UpdateTaskReplyValidationError is the validation error returned by
// UpdateTaskReply.Validate if the designated constraints aren't met.
type UpdateTaskReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTaskReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTaskReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTaskReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTaskReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTaskReplyValidationError) ErrorName() string { return "UpdateTaskReplyValidationError" }

// Error satisfies the builtin error interface
func (e UpdateTaskReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTaskReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTaskReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTaskReplyValidationError{}

// Validate checks the field values on OverwriteTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OverwriteTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OverwriteTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OverwriteTaskRequestMultiError, or nil if none found.
func (m *OverwriteTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OverwriteTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Status

	// no validation rules for Name

	// no validation rules for Comment

	if len(errors) > 0 {
		return OverwriteTaskRequestMultiError(errors)
	}

	return nil
}

// OverwriteTaskRequestMultiError is an error wrapping multiple validation
// errors returned by OverwriteTaskRequest.ValidateAll() if the designated
// constraints aren't met.
type OverwriteTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OverwriteTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OverwriteTaskRequestMultiError) AllErrors() []error { return m }

// OverwriteTaskRequestValidationError is the validation error returned by
// OverwriteTaskRequest.Validate if the designated constraints aren't met.
type OverwriteTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OverwriteTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OverwriteTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OverwriteTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OverwriteTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OverwriteTaskRequestValidationError) ErrorName() string {
	return "OverwriteTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OverwriteTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOverwriteTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OverwriteTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OverwriteTaskRequestValidationError{}

// Validate checks the field values on OverwriteTaskReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OverwriteTaskReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OverwriteTaskReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OverwriteTaskReplyMultiError, or nil if none found.
func (m *OverwriteTaskReply) ValidateAll() error {
	return m.validate(true)
}

func (m *OverwriteTaskReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RowsAffected

	if len(errors) > 0 {
		return OverwriteTaskReplyMultiError(errors)
	}

	return nil
}

// OverwriteTaskReplyMultiError is an error wrapping multiple validation errors
// returned by OverwriteTaskReply.ValidateAll() if the designated constraints
// aren't met.
type OverwriteTaskReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OverwriteTaskReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OverwriteTaskReplyMultiError) AllErrors() []error { return m }

// OverwriteTaskReplyValidationError is the validation error returned by
// OverwriteTaskReply.Validate if the designated constraints aren't met.
type OverwriteTaskReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OverwriteTaskReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OverwriteTaskReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OverwriteTaskReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OverwriteTaskReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OverwriteTaskReplyValidationError) ErrorName() string {
	return "OverwriteTaskReplyValidationError"
}

// Error satisfies the builtin error interface
func (e OverwriteTaskReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOverwriteTaskReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OverwriteTaskReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OverwriteTaskReplyValidationError{}

// Validate checks the field values on DeleteTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTaskRequestMultiError, or nil if none found.
func (m *DeleteTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteTaskRequestMultiError(errors)
	}

	return nil
}

// DeleteTaskRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTaskRequestMultiError) AllErrors() []error { return m }

// DeleteTaskRequestValidationError is the validation error returned by
// DeleteTaskRequest.Validate if the designated constraints aren't met.
type DeleteTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTaskRequestValidationError) ErrorName() string {
	return "DeleteTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTaskRequestValidationError{}

// Validate checks the field values on DeleteTaskReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteTaskReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTaskReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTaskReplyMultiError, or nil if none found.
func (m *DeleteTaskReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTaskReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RowsAffected

	if len(errors) > 0 {
		return DeleteTaskReplyMultiError(errors)
	}

	return nil
}

// DeleteTaskReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteTaskReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteTaskReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTaskReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTaskReplyMultiError) AllErrors() []error { return m }

// DeleteTaskReplyValidationError is the validation error returned by
// DeleteTaskReply.Validate if the designated constraints aren't met.
type DeleteTaskReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTaskReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTaskReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTaskReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTaskReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTaskReplyValidationError) ErrorName() string { return "DeleteTaskReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteTaskReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTaskReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTaskReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTaskReplyValidationError{}

// Validate checks the field values on GetTaskRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTaskRequestMultiError,
// or nil if none found.
func (m *GetTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetTaskRequestMultiError(errors)
	}

	return nil
}

// GetTaskRequestMultiError is an error wrapping multiple validation errors
// returned by GetTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskRequestMultiError) AllErrors() []error { return m }

// GetTaskRequestValidationError is the validation error returned by
// GetTaskRequest.Validate if the designated constraints aren't met.
type GetTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskRequestValidationError) ErrorName() string { return "GetTaskRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskRequestValidationError{}

// Validate checks the field values on GetTaskReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTaskReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTaskReplyMultiError, or
// nil if none found.
func (m *GetTaskReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Status

	// no validation rules for Name

	// no validation rules for Comment

	// no validation rules for UpdatedBy

	// no validation rules for CreatedBy

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return GetTaskReplyMultiError(errors)
	}

	return nil
}

// GetTaskReplyMultiError is an error wrapping multiple validation errors
// returned by GetTaskReply.ValidateAll() if the designated constraints aren't met.
type GetTaskReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskReplyMultiError) AllErrors() []error { return m }

// GetTaskReplyValidationError is the validation error returned by
// GetTaskReply.Validate if the designated constraints aren't met.
type GetTaskReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskReplyValidationError) ErrorName() string { return "GetTaskReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetTaskReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskReplyValidationError{}

// Validate checks the field values on SingleTask with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SingleTask) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SingleTask with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SingleTaskMultiError, or
// nil if none found.
func (m *SingleTask) ValidateAll() error {
	return m.validate(true)
}

func (m *SingleTask) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Status

	// no validation rules for Name

	// no validation rules for Comment

	// no validation rules for UpdatedBy

	// no validation rules for CreatedBy

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return SingleTaskMultiError(errors)
	}

	return nil
}

// SingleTaskMultiError is an error wrapping multiple validation errors
// returned by SingleTask.ValidateAll() if the designated constraints aren't met.
type SingleTaskMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SingleTaskMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SingleTaskMultiError) AllErrors() []error { return m }

// SingleTaskValidationError is the validation error returned by
// SingleTask.Validate if the designated constraints aren't met.
type SingleTaskValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SingleTaskValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SingleTaskValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SingleTaskValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SingleTaskValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SingleTaskValidationError) ErrorName() string { return "SingleTaskValidationError" }

// Error satisfies the builtin error interface
func (e SingleTaskValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSingleTask.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SingleTaskValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SingleTaskValidationError{}

// Validate checks the field values on GetTasksRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTasksRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTasksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTasksRequestMultiError, or nil if none found.
func (m *GetTasksRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTasksRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetTasksRequestMultiError(errors)
	}

	return nil
}

// GetTasksRequestMultiError is an error wrapping multiple validation errors
// returned by GetTasksRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTasksRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTasksRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTasksRequestMultiError) AllErrors() []error { return m }

// GetTasksRequestValidationError is the validation error returned by
// GetTasksRequest.Validate if the designated constraints aren't met.
type GetTasksRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTasksRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTasksRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTasksRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTasksRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTasksRequestValidationError) ErrorName() string { return "GetTasksRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTasksRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTasksRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTasksRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTasksRequestValidationError{}

// Validate checks the field values on GetTasksReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTasksReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTasksReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTasksReplyMultiError, or
// nil if none found.
func (m *GetTasksReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTasksReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTasks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTasksReplyValidationError{
						field:  fmt.Sprintf("Tasks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTasksReplyValidationError{
						field:  fmt.Sprintf("Tasks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTasksReplyValidationError{
					field:  fmt.Sprintf("Tasks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetTasksReplyMultiError(errors)
	}

	return nil
}

// GetTasksReplyMultiError is an error wrapping multiple validation errors
// returned by GetTasksReply.ValidateAll() if the designated constraints
// aren't met.
type GetTasksReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTasksReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTasksReplyMultiError) AllErrors() []error { return m }

// GetTasksReplyValidationError is the validation error returned by
// GetTasksReply.Validate if the designated constraints aren't met.
type GetTasksReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTasksReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTasksReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTasksReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTasksReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTasksReplyValidationError) ErrorName() string { return "GetTasksReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetTasksReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTasksReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTasksReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTasksReplyValidationError{}
