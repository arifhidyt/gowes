package respond

import "net/http"

// BadRequest returns a 400 Bad Request JSON response
func (r *Response) BadRequest(v interface{}) {
	r.writeResponse(http.StatusBadRequest, v)
}

// Unauthorized returns a 401 Unauthorized JSON response
func (r *Response) Unauthorized(v interface{}) {
	r.writeResponse(http.StatusUnauthorized, v)
}

// PaymentRequired returns a 402 Payment Required JSON response
func (r *Response) PaymentRequired(v interface{}) {
	r.writeResponse(http.StatusPaymentRequired, v)
}

// Forbidden returns a 403 Forbidden JSON response
func (r *Response) Forbidden(v interface{}) {
	r.writeResponse(http.StatusForbidden, v)
}

// NotFound returns a 404 Not Found JSON response
func (r *Response) NotFound(v interface{}) {
	r.writeResponse(http.StatusNotFound, v)
}

// MethodNotAllowed returns a 405 Method Not Allowed JSON response
func (r *Response) MethodNotAllowed(v interface{}) {
	r.writeResponse(http.StatusMethodNotAllowed, v)
}

// Conflict returns a 409 Conflict JSON response
func (r *Response) Conflict(v interface{}) {
	r.writeResponse(http.StatusConflict, v)
}

// LengthRequired returns a 411 Length Required JSON response
func (r *Response) LengthRequired(v interface{}) {
	r.writeResponse(http.StatusLengthRequired, v)
}

// UnprocessableEntity returns a 422 Unprocessable Entity JSON response
func (r *Response) UnprocessableEntity(v interface{}) {
	r.writeResponse(http.StatusUnprocessableEntity, v)
}

// InternalServerError returns a 500 Internal Server Error JSON response
func (r *Response) InternalServerError(v interface{}) {
	r.writeResponse(http.StatusInternalServerError, v)
}

// NotImplemented returns a 501 Not Implemented JSON response
func (r *Response) NotImplemented(v interface{}) {
	r.writeResponse(http.StatusNotImplemented, v)
}

// BadGateway returns a 502 Bad Gateway JSON response
func (r *Response) BadGateway(v interface{}) {
	r.writeResponse(http.StatusBadGateway, v)
}

// ServiceUnavailable returns a 503 Service Unavailable JSON response
func (r *Response) ServiceUnavailable(v interface{}) {
	r.writeResponse(http.StatusServiceUnavailable, v)
}

// GatewayTimeout returns a 504 Gateway Timeout JSON response
func (r *Response) GatewayTimeout(v interface{}) {
	r.writeResponse(http.StatusGatewayTimeout, v)
}