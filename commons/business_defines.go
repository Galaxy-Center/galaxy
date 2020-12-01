package commons

// Code defines for WebAPIResponse.
type Code int

const (
	// OK everything is ok.
	OK = iota
	// DBOperationAbnormal occurred db expection.
	DBOperationAbnormal
	// NotFound 404.
	NotFound
	// BadRequest
	BadRequest
)

// Status kk
type Status struct {
	Code   Code
	Reason string
}

func build(c Code, r string) Status {
	return Status{
		Code:   c,
		Reason: r,
	}
}

var (
	// StatusOK -> OK
	StatusOK = build(OK, "ok")
	// StatusDBOperationAbnormal -> DBOperationAbnormal
	StatusDBOperationAbnormal = build(DBOperationAbnormal, "database operation abnormal.")
	// StatusNotFound -> NotFound
	StatusNotFound = build(NotFound, "resource not found.")
	// StatusBadRequest -> BadRequest
	StatusBadRequest = build(BadRequest, "invalid request params.")
)

// WebAPIResponse wrapper of anything from bi.
type WebAPIResponse struct {
	Code    Code
	Message string
	Data    interface{}
}

// Success returns perfact rerults.
func Success(data interface{}) WebAPIResponse {
	return WebAPIResponse{
		Code:    StatusOK.Code,
		Message: StatusOK.Reason,
		Data:    data,
	}
}

// Error returns customer results.
func Error(s Status) WebAPIResponse {
	return WebAPIResponse{
		Code:    s.Code,
		Message: s.Reason,
		Data:    nil,
	}
}

// ErrorWithMessage returns customer results.
func ErrorWithMessage(c Code, m string) WebAPIResponse {
	return WebAPIResponse{
		Code:    c,
		Message: m,
		Data:    nil,
	}
}
