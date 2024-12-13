package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"atmail"
)

func TestCreateUserOk(t *testing.T) {
	var b bytes.Buffer

	if err := json.NewEncoder(&b).Encode(createUserInload{
		Username: "johndoe",
		Email:    "john@doe.com",
		Age:      42,
	}); err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/users", &b)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	s := fakeStore{newUserId: 1234}

	want := createUserOutload{
		Id:       1234,
		Username: "johndoe",
		Email:    "john@doe.com",
		Age:      42,
	}

	got, err := createUser(s, rr, req)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestCreateUserNotOk(t *testing.T) {
	store := fakeStore{
		existingUser: atmail.User{
			Username: "existinguser",
			Email:    "existing@email.com",
		},
	}

	for name, tc := range map[string]struct {
		inload string
		want   outload
	}{
		"invalid json": {
			inload: `this invalid json`,
			want: errorOutload{
				Code:  http.StatusBadRequest,
				Error: "invalid json!",
			},
		},
		"blank username": {
			inload: `{}`,
			want: errorOutload{
				Code:  http.StatusBadRequest,
				Error: "username cannot be blank!",
			},
		},
		"blank email": {
			inload: `{ "username": "some-valid-username" }`,
			want: errorOutload{
				Code:  http.StatusBadRequest,
				Error: "email is invalid!",
			},
		},
		"invalid age (negative)": {
			inload: `{ "username": "some-valid-username", "email": "valid@email.com", "age": -1 }`,
			want: errorOutload{
				Code:  http.StatusBadRequest,
				Error: "invalid json!",
			},
		},
		"username or email already exists": {
			inload: `{ "username": "existinguser", "email": "existing@email.com", "age": 1 }`,
			want: errorOutload{
				Code:  http.StatusBadRequest,
				Error: "username/email already exists!",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(tc.inload)))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			got, err := createUser(store, rr, req)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want %v; got %v", tc.want, got)
			}
		})
	}
}

func TestGetUserOk(t *testing.T) {
	s := fakeStore{
		existingUser: atmail.User{
			Id:       4321,
			Username: "janedoe",
			Email:    "jane@doe.com",
			Age:      24,
		},
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("/users/%d", s.existingUser.Id), nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", strconv.FormatInt(s.existingUser.Id, 10))

	rr := httptest.NewRecorder()

	want := getUserOutload{
		Id:       4321,
		Username: "janedoe",
		Email:    "jane@doe.com",
		Age:      24,
	}

	got, err := getUser(s, rr, req)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestGetUserNotOk(t *testing.T) {
	for name, tc := range map[string]struct {
		want  outload
		store fakeStore
		id    string
	}{
		"user does not exist": {
			want: errorOutload{
				Code:  http.StatusBadRequest,
				Error: "user does not exist!",
			},
			store: fakeStore{},
			id:    "999",
		},
		"invalid id format": {
			want: errorOutload{
				Code:  http.StatusBadRequest,
				Error: "user does not exist!",
			},
			store: fakeStore{},
			id:    "some-invalid-id",
		},
	} {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("/users/%s", tc.id), nil)
			if err != nil {
				t.Fatal(err)
			}

			req.SetPathValue("id", tc.id)

			rr := httptest.NewRecorder()

			got, err := getUser(tc.store, rr, req)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want %v; got %v", tc.want, got)
			}
		})
	}
}

func TestUpdateUserOk(t *testing.T) {
	store := fakeStore{
		existingUser: atmail.User{
			Id:       1234,
			Username: "some-username",
			Email:    "valid@email.com",
			Age:      64,
		},
	}

	for name, tc := range map[string]struct {
		inload string
		id     string
		want   outload
	}{
		"username only provided": {
			inload: `{ "username": "something-else" }`,
			want: updateUserOutload{
				Id:       1234,
				Username: "something-else",
				Email:    "valid@email.com",
				Age:      64,
			},
		},
		"email only provided": {
			inload: `{ "email": "foo@bar.com" }`,
			want: updateUserOutload{
				Id:       1234,
				Username: "some-username",
				Email:    "foo@bar.com",
				Age:      64,
			},
		},
		"age only provided ": {
			inload: `{ "age": 123 }`,
			want: updateUserOutload{
				Id:       1234,
				Username: "some-username",
				Email:    "valid@email.com",
				Age:      123,
			},
		},
		"all provided": {
			inload: `{ "username": "bandit", "email": "new@email.com", "age": 321 }`,
			want: updateUserOutload{
				Id:       1234,
				Username: "bandit",
				Email:    "new@email.com",
				Age:      321,
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("PUT", "/users/1234", bytes.NewBuffer([]byte(tc.inload)))
			if err != nil {
				t.Fatal(err)
			}

			req.SetPathValue("id", "1234")

			rr := httptest.NewRecorder()

			got, err := updateUser(store, rr, req)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want %v; got %v", tc.want, got)
			}
		})
	}
}

func TestUpdateUserNotOk(t *testing.T) {
	store := fakeStore{
		existingUser: atmail.User{
			Id:       1234,
			Username: "some-username",
			Email:    "valid@email.com",
			Age:      64,
		},
	}

	for name, tc := range map[string]struct {
		inload string
		id     string
		want   outload
	}{
		"empty body": {
			inload: `{ "username": "" }`,
			want: errorOutload{
				Code:  http.StatusBadRequest,
				Error: "username cannot be blank!",
			},
		},
		"blank username": {
			inload: `{ "username": "" }`,
			want: errorOutload{
				Code:  http.StatusBadRequest,
				Error: "username cannot be blank!",
			},
		},
		"invalid email": {
			inload: `{ "email": "invalid-email"}`,
			want: errorOutload{
				Code:  http.StatusBadRequest,
				Error: "email is invalid!",
			},
		},
		"invalid age": {
			inload: `{ "age": 0 }`,
			want: errorOutload{
				Code:  http.StatusBadRequest,
				Error: "age is invalid!",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("PUT", "/users/1234", bytes.NewBuffer([]byte(tc.inload)))
			if err != nil {
				t.Fatal(err)
			}

			req.SetPathValue("id", "1234")

			rr := httptest.NewRecorder()

			got, err := updateUser(store, rr, req)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want %v; got %v", tc.want, got)
			}
		})
	}
}

func TestDeleteUserOk(t *testing.T) {
	store := fakeStore{}

	for name, tc := range map[string]struct {
		id   string
		want outload
	}{
		"ok": {
			id:   "12345",
			want: messageOutload{"successfully deleted user 12345!"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("/users/%s", tc.id), nil)
			if err != nil {
				t.Fatal(err)
			}

			req.SetPathValue("id", tc.id)

			rr := httptest.NewRecorder()

			got, err := deleteUser(store, rr, req)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want %v; got %v", tc.want, got)
			}
		})
	}
}
