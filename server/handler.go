package server

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"atmail"
	"atmail/server/roles"
)

type handlerFunc func(atmail.Store, http.ResponseWriter, *http.Request) (outload, error)

type handler struct {
	store atmail.Store
	roles roles.Role
	fn    handlerFunc
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// set headers
	w.Header().Add("Content-Type", "application/json")

	token, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Basic ")
	if !ok {
		fail(w, "unauthorized!", http.StatusUnauthorized)
		return
	}

	bs, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		fail(w, "internal server error!", http.StatusInternalServerError)
		return
	}

	ss := strings.Split(string(bs), ":")

	if len(ss) != 2 {
		fail(w, "unauthorized!", http.StatusUnauthorized)
		return
	}

	user, password := ss[0], ss[1]

	role, err := h.store.GetRole(user, password)
	if err != nil {
		if !errors.Is(err, atmail.ErrAdminNone) {
			fail(w, "internal server error!", http.StatusInternalServerError)
			return
		}

		fail(w, "unauthorized!", http.StatusUnauthorized)
		return
	}

	if !roles.IsAuthorized(h.roles, role) {
		fail(w, "unauthorized!", http.StatusUnauthorized)
		return
	}

	// get outload
	outload, err := h.fn(h.store, w, r)
	if err != nil {
		fail(w, "internal server error!", http.StatusInternalServerError)
		return
	}

	// write json
	write(w, outload)
}

func write(w http.ResponseWriter, o outload) {
	w.WriteHeader(o.code())

	encoder := json.NewEncoder(w)

	encoder.SetIndent("", "\t")

	if err := encoder.Encode(&o); err != nil {
		fail(w, "internal server error!", http.StatusInternalServerError)
		return
	}
}

func fail(w http.ResponseWriter, message string, code int) {
	write(w, errorOutload{code, message})
}
