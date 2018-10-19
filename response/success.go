package respond

import (
	"net/http"
)

// Ok returns a 200 OK JSON response
func (r *Response) Ok(v interface{}) {
	r.writeResponse(http.StatusOK, v)
}

// Created returns a 201 Created JSON response
func (r *Response) Created(v interface{}) {
	r.writeResponse(http.StatusCreated, v)
}

// Accepted returns a 202 Accepted JSON response
func (r *Response) Accepted(v interface{}) {
	r.writeResponse(http.StatusAccepted, v)
}

// NoContent returns a 204 No Content JSON response
func (r *Response) NoContent() {
	r.writeResponse(http.StatusNoContent, nil)
}