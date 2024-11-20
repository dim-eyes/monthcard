package encoder

import (
	"encoding/json"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
)

type ErrorResponse struct {
	Code     int         `json:"code"`
	Reason   string      `json:"reason"`
	Message  string      `json:"message"`
	Metadata interface{} `json:"metadata"`
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	var code int
	var message string
	var metadata interface{}
	if kratosErr, ok := err.(*errors.Error); ok {
		code = int(kratosErr.Code)
		message = kratosErr.Message
	} else {
		code = http.StatusNetworkAuthenticationRequired
		message = http.StatusText(code)
	}

	response := ErrorResponse{
		Code:     code,
		Message:  message,
		Metadata: metadata,
	}
	json.NewEncoder(w).Encode(response)
}
