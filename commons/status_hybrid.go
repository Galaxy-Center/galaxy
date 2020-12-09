package commons

import (
	"errors"
	"fmt"
	"net/http"
)

// Adds definition status code
const (
	// DBOperationAbnormal occurred db expection.
	DBOperationAbnormal = http.StatusInternalServerError
)

// Error customer.
type Error struct {
	Code  int
	Error error
}

// Format returns string of curr.
func (r *Error) Format() string {
	return fmt.Sprintf("status %d: err %v", r.Code, r.Error)
}

var (
	// OK -> 0
	OK = &Error{Code: 0, Error: nil}
	// StatusDBOperationAbnormal -> DBOperationAbnormal
	StatusDBOperationAbnormal = &Error{Code: DBOperationAbnormal, Error: errors.New("database operation abnormal")}
)

// WebAPIResponse wrapper of anything from bi.
type WebAPIResponse struct {
	Code    int
	Message string
	Data    interface{}
}

// Success returns perfact rerults.
func Success(data interface{}) WebAPIResponse {
	return WebAPIResponse{
		Code:    0,
		Message: "ok",
		Data:    data,
	}
}

// ErrorWithCustomer returns customer results.
func ErrorWithCustomer(r Error) WebAPIResponse {
	return WebAPIResponse{
		Code:    r.Code,
		Message: r.Format(),
		Data:    nil,
	}
}

// ErrorWithMessage returns customer results.
func ErrorWithMessage(m string) WebAPIResponse {
	return WebAPIResponse{
		Code:    1,
		Message: m,
		Data:    nil,
	}
}
