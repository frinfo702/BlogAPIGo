package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/frinfo702/MyApi/api/middlewares"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	// return appropriate http response based on error types

	var appErr *MyAppError
	if !errors.As(err, &appErr) {
		// if err's error tree does not contain MyAppError, return "Unknown" error
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := middlewares.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int
	switch appErr.ErrCode {
	case EmptyData:
		statusCode = http.StatusNotFound // 404
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest // 400
	default:
		statusCode = http.StatusInternalServerError // 500
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
