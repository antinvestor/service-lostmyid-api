// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: mylostid.proto

package mylostidv1

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
)

// Validate checks the field values on CollectibleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CollectibleRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := len(m.GetFileId()); l < 1 || l > 5 {
		return CollectibleRequestValidationError{
			field:  "FileId",
			reason: "value must contain between 1 and 5 items, inclusive",
		}
	}

	for idx, item := range m.GetFileId() {
		_, _ = idx, item

		if l := utf8.RuneCountInString(item); l < 3 || l > 40 {
			return CollectibleRequestValidationError{
				field:  fmt.Sprintf("FileId[%v]", idx),
				reason: "value length must be between 3 and 40 runes, inclusive",
			}
		}

		if !_CollectibleRequest_FileId_Pattern.MatchString(item) {
			return CollectibleRequestValidationError{
				field:  fmt.Sprintf("FileId[%v]", idx),
				reason: "value does not match regex pattern \"[0-9a-z_-]{3,20}\"",
			}
		}

	}

	// no validation rules for Properties

	return nil
}

// CollectibleRequestValidationError is the validation error returned by
// CollectibleRequest.Validate if the designated constraints aren't met.
type CollectibleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CollectibleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CollectibleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CollectibleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CollectibleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CollectibleRequestValidationError) ErrorName() string {
	return "CollectibleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CollectibleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCollectibleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CollectibleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CollectibleRequestValidationError{}

var _CollectibleRequest_FileId_Pattern = regexp.MustCompile("[0-9a-z_-]{3,20}")

// Validate checks the field values on CollectibleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CollectibleResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Properties

	// no validation rules for State

	return nil
}

// CollectibleResponseValidationError is the validation error returned by
// CollectibleResponse.Validate if the designated constraints aren't met.
type CollectibleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CollectibleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CollectibleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CollectibleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CollectibleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CollectibleResponseValidationError) ErrorName() string {
	return "CollectibleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CollectibleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCollectibleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CollectibleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CollectibleResponseValidationError{}

// Validate checks the field values on SearchRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SearchRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetId()); l < 3 || l > 40 {
		return SearchRequestValidationError{
			field:  "Id",
			reason: "value length must be between 3 and 40 runes, inclusive",
		}
	}

	if !_SearchRequest_Id_Pattern.MatchString(m.GetId()) {
		return SearchRequestValidationError{
			field:  "Id",
			reason: "value does not match regex pattern \"[0-9a-z_-]{3,40}\"",
		}
	}

	// no validation rules for Query

	// no validation rules for Properties

	return nil
}

// SearchRequestValidationError is the validation error returned by
// SearchRequest.Validate if the designated constraints aren't met.
type SearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchRequestValidationError) ErrorName() string { return "SearchRequestValidationError" }

// Error satisfies the builtin error interface
func (e SearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchRequestValidationError{}

var _SearchRequest_Id_Pattern = regexp.MustCompile("[0-9a-z_-]{3,40}")

// Validate checks the field values on SearchResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SearchResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Query

	// no validation rules for Properties

	// no validation rules for State

	return nil
}

// SearchResponseValidationError is the validation error returned by
// SearchResponse.Validate if the designated constraints aren't met.
type SearchResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchResponseValidationError) ErrorName() string { return "SearchResponseValidationError" }

// Error satisfies the builtin error interface
func (e SearchResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchResponseValidationError{}

// Validate checks the field values on ProgressRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ProgressRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetId()); l < 3 || l > 40 {
		return ProgressRequestValidationError{
			field:  "Id",
			reason: "value length must be between 3 and 40 runes, inclusive",
		}
	}

	if !_ProgressRequest_Id_Pattern.MatchString(m.GetId()) {
		return ProgressRequestValidationError{
			field:  "Id",
			reason: "value does not match regex pattern \"[0-9a-z_-]{3,40}\"",
		}
	}

	if _, ok := ITEMType_name[int32(m.GetType())]; !ok {
		return ProgressRequestValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// ProgressRequestValidationError is the validation error returned by
// ProgressRequest.Validate if the designated constraints aren't met.
type ProgressRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProgressRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProgressRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProgressRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProgressRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProgressRequestValidationError) ErrorName() string { return "ProgressRequestValidationError" }

// Error satisfies the builtin error interface
func (e ProgressRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProgressRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProgressRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProgressRequestValidationError{}

var _ProgressRequest_Id_Pattern = regexp.MustCompile("[0-9a-z_-]{3,40}")

// Validate checks the field values on TransactionItem with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TransactionItem) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Datetime

	// no validation rules for Credit

	if v, ok := interface{}(m.GetAmount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TransactionItemValidationError{
				field:  "Amount",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Description

	return nil
}

// TransactionItemValidationError is the validation error returned by
// TransactionItem.Validate if the designated constraints aren't met.
type TransactionItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactionItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactionItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactionItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactionItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactionItemValidationError) ErrorName() string { return "TransactionItemValidationError" }

// Error satisfies the builtin error interface
func (e TransactionItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransactionItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactionItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactionItemValidationError{}

// Validate checks the field values on ProgressItem with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ProgressItem) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Datetime

	// no validation rules for State

	// no validation rules for Description

	return nil
}

// ProgressItemValidationError is the validation error returned by
// ProgressItem.Validate if the designated constraints aren't met.
type ProgressItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProgressItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProgressItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProgressItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProgressItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProgressItemValidationError) ErrorName() string { return "ProgressItemValidationError" }

// Error satisfies the builtin error interface
func (e ProgressItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProgressItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProgressItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProgressItemValidationError{}

// Validate checks the field values on ProgressResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ProgressResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Type

	// no validation rules for Properties

	// no validation rules for State

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProgressResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetTransactions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProgressResponseValidationError{
					field:  fmt.Sprintf("Transactions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ProgressResponseValidationError is the validation error returned by
// ProgressResponse.Validate if the designated constraints aren't met.
type ProgressResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProgressResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProgressResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProgressResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProgressResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProgressResponseValidationError) ErrorName() string { return "ProgressResponseValidationError" }

// Error satisfies the builtin error interface
func (e ProgressResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProgressResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProgressResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProgressResponseValidationError{}

// Validate checks the field values on Pagination with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Pagination) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Count

	// no validation rules for Page

	return nil
}

// PaginationValidationError is the validation error returned by
// Pagination.Validate if the designated constraints aren't met.
type PaginationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginationValidationError) ErrorName() string { return "PaginationValidationError" }

// Error satisfies the builtin error interface
func (e PaginationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPagination.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginationValidationError{}

// Validate checks the field values on RangeSpanRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RangeSpanRequest) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Range.(type) {

	case *RangeSpanRequest_Pagination:

		if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RangeSpanRequestValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *RangeSpanRequest_Interval:

		if v, ok := interface{}(m.GetInterval()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RangeSpanRequestValidationError{
					field:  "Interval",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return RangeSpanRequestValidationError{
			field:  "Range",
			reason: "value is required",
		}

	}

	return nil
}

// RangeSpanRequestValidationError is the validation error returned by
// RangeSpanRequest.Validate if the designated constraints aren't met.
type RangeSpanRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RangeSpanRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RangeSpanRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RangeSpanRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RangeSpanRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RangeSpanRequestValidationError) ErrorName() string { return "RangeSpanRequestValidationError" }

// Error satisfies the builtin error interface
func (e RangeSpanRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRangeSpanRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RangeSpanRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RangeSpanRequestValidationError{}
