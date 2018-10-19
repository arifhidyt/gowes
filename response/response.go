// HTTP status code handle and feedback message API
package respond

import (
	"encoding/json"
	"net/http"
)

// Response is the HTTP response
type Response struct {
	Writer  http.ResponseWriter
	Headers map[string]string
}

// NewResponse creates and returns a new response
func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		Writer: w,
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
	}
}

// WriteResponse writes the HTTP response status, headers and body
func (r *Response) writeResponse(code int, v interface{}) error {
	if len(r.Headers) > 0 {
		r.writeHeaders()
	}

	r.writeStatusCode(code)

	if v != nil {
		body, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		if _, err := r.Writer.Write(body); err != nil {
			panic(err)
		}
	}
	return nil
}

func (r *Response) writeHeaders() {
	for key, value := range r.Headers {
		r.Writer.Header().Set(key, value)
	}
}

func (r *Response) writeStatusCode(code int) {
	r.Writer.WriteHeader(code)
}