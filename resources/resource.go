package resources

// Code defines for WebAPIResponse.
type Code int

const (
	// OK everything is ok.
	OK Code = iota
	// NOTFOUND 404.
	NOTFOUND
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
		Code:    OK,
		Message: "",
		Data:    data,
	}
}

// Error returns customer results.
func Error(code Code, message string) WebAPIResponse {
	return WebAPIResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
