package server

import (
	"net/http"

	"atmail"
	"atmail/server/roles"
)

func New(store atmail.Store) *http.ServeMux {
	// convenience closure
	toHandler := func(fn handlerFunc, roles roles.Role) http.Handler {
		return handler{store, roles, fn}
	}

	mux := http.NewServeMux()

	mux.Handle("GET /users/{id}", toHandler(getUser,
		roles.Bingo|roles.Bluey|roles.Chilli|roles.Bandit,
	))

	mux.Handle("POST /users", toHandler(createUser,
		roles.Bluey|roles.Chilli|roles.Bandit,
	))

	mux.Handle("PUT /users/{id}", toHandler(updateUser,
		roles.Chilli|roles.Bandit,
	))

	mux.Handle("DELETE /users/{id}", toHandler(deleteUser,
		roles.Bandit,
	))

	return mux
}
