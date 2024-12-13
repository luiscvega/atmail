package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"atmail"
	"atmail/server/roles"
)

type fakeOutload struct{}

func (o fakeOutload) code() int {
	return http.StatusOK
}

func TestHandlerServeHTTPOk(t *testing.T) {
	h := handler{
		store: fakeStore{
			existingAdminUser:     "foo",
			existingAdminPassword: "bar",
			existingAdminRole:     roles.Chilli,
		},
		roles: roles.Chilli | roles.Bandit,
		fn: func(atmail.Store, http.ResponseWriter, *http.Request) (outload, error) {
			return fakeOutload{}, nil
		},
	}

	req := http.Request{Header: http.Header{}}

	req.SetBasicAuth("foo", "bar")

	rr := httptest.NewRecorder()

	h.ServeHTTP(rr, &req)

	if got := rr.Code; got != http.StatusOK {
		t.Errorf("want %v; got %v", http.StatusOK, got)
	}
}

func TestHandlerServeHTTPUnauthorized(t *testing.T) {
	h := handler{
		store: fakeStore{
			existingAdminUser:     "foo",
			existingAdminPassword: "bar",
			existingAdminRole:     roles.Chilli,
		},
		roles: roles.Bingo | roles.Bandit,
		fn: func(atmail.Store, http.ResponseWriter, *http.Request) (outload, error) {
			return nil, nil
		},
	}

	for name, tc := range map[string]struct {
		username string
		password string
	}{
		"admin does not exist": {
			username: "baz",
			password: "qux",
		},
		"role not allowed": {
			username: "foo",
			password: "bar",
		},
	} {
		t.Run(name, func(t *testing.T) {
			req := http.Request{Header: http.Header{}}

			req.SetBasicAuth(tc.username, tc.password)

			rr := httptest.NewRecorder()

			h.ServeHTTP(rr, &req)

			if got := rr.Code; got != http.StatusUnauthorized {
				t.Errorf("want %v; got %v", http.StatusUnauthorized, got)
			}

		})
	}
}
