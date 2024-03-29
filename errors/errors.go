package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

// ErrorType is the type of an error
type ErrorType uint

const (
	// NoType error
	NoType ErrorType = iota
	// BadTransaction error
	BadTransaction
	// FailedTransaction error
	FailedTransaction
	// FailedConnection error
	FailedConnection
    // FailedKeystore error
	FailedKeystore
	// FailedReadFile error
	FailedReadFile
	// FailedReadEnv error
	FailedReadEnv
	// FailedKeyConfig error
	FailedKeyConfig
	// FailedConfigTransaction error
	FailedConfigTransaction
	// FailedContract error
	FailedContract
	// CallBlockchainFailed error
	CallBlockchainFailed
	// MalformedRawTransaction error
	MalformedRawTransaction
	//InvalidAddress error
	InvalidAddress
)	

type customError struct {
	errorType ErrorType
	originalError error
	context errorContext
	errorCode int
}

type errorContext struct {
	Field string
	Message string
}

// New creates a new customError
func (errorType ErrorType) New(msg string, code int) error {
	return customError{errorType: errorType, originalError: errors.New(msg), errorCode: code}
}

// New creates a new customError with formatted message
func (errorType ErrorType) Newf(msg string, args ...interface{}) error {
	return customError{errorType: errorType, originalError: fmt.Errorf(msg, args...)}
}

// Wrap creates a new wrapped error
func (errorType ErrorType) Wrap(err error, msg string, code int) error {
	return errorType.Wrapf(err, msg, code)
}

// Wrap creates a new wrapped error with formatted message
func (errorType ErrorType) Wrapf(err error, msg string, code int, args ...interface{}) error {
	return customError{errorType: errorType, originalError: errors.Wrapf(err, msg, args...), errorCode: code}
}

// Error returns the mssage of a customError
func (error customError) Error() string {
	return error.originalError.Error()
}

//ErrorCode returns code
func (error customError) ErrorCode() int {
	return error.errorCode
}


// New creates a no type error
func New(msg string, code int) error {
	return customError{errorType: NoType, originalError: errors.New(msg), errorCode: code}
}

// Newf creates a no type error with formatted message
func Newf(msg string, code int, args ...interface{}) error {
	return customError{errorType: NoType, originalError: errors.New(fmt.Sprintf(msg, args...)), errorCode: code}
}

// Wrap an error with a string
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Cause gives the original error
func Cause(err error) error {
	return errors.Cause(err)
}

// Wrapf an error with format string
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType: customErr.errorType,
			originalError: wrappedError,
			context: customErr.context,
		}
	}

	return customError{errorType: NoType, originalError: wrappedError}
}

// AddErrorContext adds a context to an error
func AddErrorContext(err error, field, message string) error {
	context := errorContext{Field: field, Message: message}
	if customErr, ok := err.(customError); ok {
		return customError{errorType: customErr.errorType, originalError: customErr.originalError, context: context}
	}

	return customError{errorType: NoType, originalError: err, context: context}
}

// GetErrorContext returns the error context
func GetErrorContext(err error) map[string]string {
	emptyContext := errorContext{}
	if customErr, ok := err.(customError); ok || customErr.context != emptyContext  {

		return map[string]string{"field": customErr.context.Field, "message": customErr.context.Message}
	}

	return nil
}

// GetType returns the error type
func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}

	return NoType
}