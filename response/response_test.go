package respond_test

import (
	"errors"
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"

	res "github.com/arifhidyt/gowes/response"
)

func newRequest(t *testing.T, method string) *http.Request {
	req, err := http.NewRequest(method, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func validateStatusCode(responseStatusCode int, expectedStatusCode int) error {
	if responseStatusCode != expectedStatusCode {
		return errors.New(fmt.Sprintf("Handler returned undefined status code: got %v wanted %v",
			responseStatusCode, expectedStatusCode))
	}
	return nil
}

func validateResponseBody(responseBody string, expectedBody string) error {
	if responseBody != expectedBody {
		return errors.New(fmt.Sprintf("Handler returned unexpected body: got %v wanted %v",
			responseBody, expectedBody))
	}
	return nil
}

func validateResponseHeader(responseHeaderValue string, expectedHeaderValue string) error {
	if responseHeaderValue != expectedHeaderValue {
		return errors.New(fmt.Sprintf("Handler returned unexpected body: got %v wanted %v",
			responseHeaderValue, expectedHeaderValue))
	}
	return nil
}

func TestContentTypeHeader(t *testing.T) {
	t.Parallel()

	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res.NewResponse(w).
			Ok(nil)
	})
	handler.ServeHTTP(rr, req)

	contentType := "application/json; charset=utf-8"
	if err := validateResponseHeader(rr.Header().Get("Content-Type"), contentType); err != nil {
		t.Fatal(err.Error())
	}
}