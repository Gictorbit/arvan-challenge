// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: wallet/v1/wallet.proto

package walletv1

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

// Validate checks the field values on MyWalletRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MyWalletRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MyWalletRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MyWalletRequestMultiError, or nil if none found.
func (m *MyWalletRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *MyWalletRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return MyWalletRequestMultiError(errors)
	}

	return nil
}

// MyWalletRequestMultiError is an error wrapping multiple validation errors
// returned by MyWalletRequest.ValidateAll() if the designated constraints
// aren't met.
type MyWalletRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MyWalletRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MyWalletRequestMultiError) AllErrors() []error { return m }

// MyWalletRequestValidationError is the validation error returned by
// MyWalletRequest.Validate if the designated constraints aren't met.
type MyWalletRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MyWalletRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MyWalletRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MyWalletRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MyWalletRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MyWalletRequestValidationError) ErrorName() string { return "MyWalletRequestValidationError" }

// Error satisfies the builtin error interface
func (e MyWalletRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMyWalletRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MyWalletRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MyWalletRequestValidationError{}

// Validate checks the field values on MyWalletResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MyWalletResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MyWalletResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MyWalletResponseMultiError, or nil if none found.
func (m *MyWalletResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MyWalletResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Balance

	for idx, item := range m.GetTransactions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MyWalletResponseValidationError{
						field:  fmt.Sprintf("Transactions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MyWalletResponseValidationError{
						field:  fmt.Sprintf("Transactions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MyWalletResponseValidationError{
					field:  fmt.Sprintf("Transactions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MyWalletResponseMultiError(errors)
	}

	return nil
}

// MyWalletResponseMultiError is an error wrapping multiple validation errors
// returned by MyWalletResponse.ValidateAll() if the designated constraints
// aren't met.
type MyWalletResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MyWalletResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MyWalletResponseMultiError) AllErrors() []error { return m }

// MyWalletResponseValidationError is the validation error returned by
// MyWalletResponse.Validate if the designated constraints aren't met.
type MyWalletResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MyWalletResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MyWalletResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MyWalletResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MyWalletResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MyWalletResponseValidationError) ErrorName() string { return "MyWalletResponseValidationError" }

// Error satisfies the builtin error interface
func (e MyWalletResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMyWalletResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MyWalletResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MyWalletResponseValidationError{}

// Validate checks the field values on Transaction with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Transaction) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Transaction with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TransactionMultiError, or
// nil if none found.
func (m *Transaction) ValidateAll() error {
	return m.validate(true)
}

func (m *Transaction) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Amount

	// no validation rules for Description

	// no validation rules for Timestamp

	if len(errors) > 0 {
		return TransactionMultiError(errors)
	}

	return nil
}

// TransactionMultiError is an error wrapping multiple validation errors
// returned by Transaction.ValidateAll() if the designated constraints aren't met.
type TransactionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TransactionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TransactionMultiError) AllErrors() []error { return m }

// TransactionValidationError is the validation error returned by
// Transaction.Validate if the designated constraints aren't met.
type TransactionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactionValidationError) ErrorName() string { return "TransactionValidationError" }

// Error satisfies the builtin error interface
func (e TransactionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransaction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactionValidationError{}
