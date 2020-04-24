package lib

import "fmt"

// Error represents error interface with custom fields.
type Error struct {
	Code    int64 `json:"code"`
	Details error `json:"message"`
}

// Error is a method for error-interface implementation.
func (e Error) Error() string {
	return fmt.Sprintf("result: %t; code: %d; message: %s", false, e.Code, e.Details.Error())
}

// NewError Returns a new Error
func NewError(codePrefix, code int64, err error) Error {
	return Error{
		// Error Codes will follow this pattern: (PREFIX) + XXX (error code)
		Code:    1000*codePrefix + code,
		Details: err,
	}
}
