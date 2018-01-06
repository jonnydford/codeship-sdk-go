package codeship

import "net/http"

const (
	defaultBaseURL = "https://api.codeship.com/v2"
	mediaType      = "application/json"
)

// Response is a DigitalOcean response. This wraps the standard http.Response returned from DigitalOcean.
type Response struct {
	*http.Response

	// Links that were returned with the response. These are parsed from
	// request body and not the header.
	Links *Links

	// Monitoring URI
	Monitor string

	Rate
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message
	Message string `json:"message"`

	// RequestID returned from the API, useful to contact support.
	RequestID string `json:"request_id"`
}
