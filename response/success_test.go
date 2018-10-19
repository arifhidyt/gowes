package respond_test

import (
	"testing"
	"net/http"
	"net/http/httptest"

	res "github.com/arifhidyt/gowes/response"
)

func TestOk(t *testing.T) {
	t.Parallel()

	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res.NewResponse(w).
			Ok(nil)
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusOK); err != nil {
		t.Fatal(err.Error())
	}

	if err := validateResponseBody(rr.Body.String(), ""); err != nil {
		t.Fatal(err.Error())
	}
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TestCreated(t *testing.T) {
	t.Parallel()

	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			{1, "arif", "arif@example.com"},
			{2, "dayat", "dayat@example.com"},
		}

		res.NewResponse(w).
			Created(users)
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusCreated); err != nil {
		t.Fatal(err.Error())
	}

	expected := `[{"id":1,"name":"arif","email":"arif@example.com"},` +
		`{"id":2,"name":"dayat","email":"dayat@example.com"}]`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}