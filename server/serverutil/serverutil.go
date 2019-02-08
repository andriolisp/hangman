package serverutil

import (
	"encoding/json"
	"errors"
	"net/http"
)

func response(code int, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Keep-Alive", "timeout=5, max=1000")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func ResponseHTMLOK(w http.ResponseWriter, r *http.Request, html []byte) {
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Keep-Alive", "timeout=5, max=1000")
	w.WriteHeader(200)

	w.Write(html)
}

// ResponseAPIOK returns a standard API success response
func ResponseAPIOK(w http.ResponseWriter, r *http.Request, data interface{}) {
	response(http.StatusOK, w, data)
}

//ResponseMethodNotAllowed return the error 405
func ResponseMethodNotAllowed(w http.ResponseWriter, r *http.Request, message string) {
	response(http.StatusMethodNotAllowed, w, errors.New("Invalid action"))
}

// ResponseNoContent returns a standard API success with no content response
func ResponseNoContent(w http.ResponseWriter, r *http.Request) {
	response(http.StatusNoContent, w, nil)
}

// ResponseAPIError returns a standard API error to the response
func ResponseAPIError(w http.ResponseWriter, r *http.Request, message string) {
	response(http.StatusBadRequest, w, errors.New(message))
}

// ResponseAPIAuthError returns a standard API auth error to the response
func ResponseAPIAuthError(w http.ResponseWriter, r *http.Request, message string) {
	response(http.StatusForbidden, w, errors.New(message))
}

// ResponseAPIAuthorizationError returns a standard API auth error to the response
func ResponseAPIAuthorizationError(w http.ResponseWriter, r *http.Request) {
	response(http.StatusForbidden, w, errors.New("Not Authorized"))
}

// ResponseAPIServiceError returns a standard API service unavailable error to the response
func ResponseAPIServiceError(w http.ResponseWriter, r *http.Request, message string) {
	response(http.StatusServiceUnavailable, w, errors.New(message))
}

// ResponseAPIValidationError returns a standard API validation error to the response
func ResponseAPIValidationError(w http.ResponseWriter, r *http.Request, message string) {
	response(http.StatusUnprocessableEntity, w, errors.New(message))
}

// ResponseAPICustomValidationError returns a standard API validation error with custom data to the response
func ResponseAPICustomValidationError(w http.ResponseWriter, r *http.Request, message string, data interface{}) {
	response(http.StatusUnprocessableEntity, w, errors.New(message))
}

// ResponseAPIFieldValidationError returns a standard API field validation error to the response
func ResponseAPIFieldValidationError(w http.ResponseWriter, r *http.Request, field string, message string) {
	response(http.StatusUnprocessableEntity, w, errors.New(field))
}

// ResponseAPINotFoundError returns a standard API not found error to the response
func ResponseAPINotFoundError(w http.ResponseWriter, r *http.Request) {
	response(http.StatusNotFound, w, errors.New("Not Found"))
}

//ResponseAPIForbiddenWithMessage returns a standard API error with the message
func ResponseAPIForbiddenWithMessage(w http.ResponseWriter, r *http.Request, message string) {
	response(http.StatusForbidden, w, errors.New(message))
}
