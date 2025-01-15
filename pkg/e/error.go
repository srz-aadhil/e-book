package e

import (
	"net/http"
	"strconv"
)

type WrapError struct {
	ErrorCode int
	Message   string
	RootCause error
}

type HttpError struct {
	Statuscode int
	Code       int
	Message    string
}

func (e *WrapError) Error() string {
	return e.RootCause.Error()
}

func NewError(errCode int, msg string, rootCause error) *WrapError {
	err := &WrapError{
		ErrorCode: errCode,
		Message:   msg,
		RootCause: rootCause,
	}

	return err
}

func NewAPIError(err error, msg string) *HttpError {
	if err == nil {
		return nil
	}

	appErr, ok := err.(*WrapError)
	if ok {
		appErr.Message = msg
	} else {
		return nil
	}
	httpError := &HttpError{
		Statuscode: GetHttpStatusCode(appErr.ErrorCode),
		Code:       appErr.ErrorCode,
		Message:    msg,
	}

	return httpError
}

// function to trim 6 digit errorcode to standard http error code
func GetHttpStatusCode(c int) int {
	//converting integer to sring
	str := strconv.Itoa(c)

	//trimming the code to 3 digit
	code := str[:3]

	//converting string back to integer type
	r, _ := strconv.Atoi(code)
	if r < 100 || r >= 600 {
		return http.StatusInternalServerError
	}

	return r

}
