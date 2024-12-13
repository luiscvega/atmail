package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"atmail"
)

type createUserInload struct {
	Username string
	Email    string
	Age      uint
}

type createUserOutload struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

func (o createUserOutload) code() int {
	return http.StatusCreated
}

func createUser(s atmail.Store, w http.ResponseWriter, r *http.Request) (outload, error) {
	i := createUserInload{}

	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		return badRequest("invalid json!"), nil
	}

	user := atmail.User{
		Username: i.Username,
		Email:    i.Email,
		Age:      i.Age,
	}

	if s, ok := atmail.ValidateUser(user); !ok {
		return badRequest(s), nil
	}

	exists, err := s.CheckUser(user.Username, user.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return badRequest("username/email already exists!"), nil
	}

	id, err := s.CreateUser(user)
	if err != nil {
		return nil, err
	}

	o := createUserOutload{
		Id:       id,
		Username: i.Username,
		Email:    i.Email,
		Age:      uint(i.Age),
	}

	return o, nil
}

type getUserOutload struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

func (o getUserOutload) code() int {
	return http.StatusOK
}

func getUser(s atmail.Store, w http.ResponseWriter, r *http.Request) (outload, error) {
	// if {id} is invalid, just return that the user does not exist
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		return badRequest("user does not exist!"), nil
	}

	user, err := s.GetUser(id)
	if err != nil {
		if !errors.Is(err, atmail.ErrUserNone) {
			return nil, err
		}

		return badRequest("user does not exist!"), nil
	}

	o := getUserOutload{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}

	return o, nil
}

type outload interface {
	code() int
}

func badRequest(message string) errorOutload {
	return errorOutload{
		Code:  http.StatusBadRequest,
		Error: message,
	}
}

type errorOutload struct {
	Code  int    `json:"-"`
	Error string `json:"error"`
}

func (e errorOutload) code() int {
	return e.Code
}

type updateUserInload struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Age      *uint   `json:"age"`
}

type updateUserOutload struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

func (o updateUserOutload) code() int {
	return http.StatusOK
}

func updateUser(s atmail.Store, w http.ResponseWriter, r *http.Request) (outload, error) {
	// if {id} is invalid, just return that the user does not exist
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		return badRequest("user does not exist!"), nil
	}

	user, err := s.GetUser(id)
	if err != nil {
		if err != atmail.ErrUserNone {
			return nil, err
		}

		return badRequest("user does not exist!"), nil
	}

	i := updateUserInload{}

	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		return nil, err
	}

	if i.Username != nil {
		user.Username = *i.Username
	}

	if i.Email != nil {
		user.Email = *i.Email
	}

	if i.Age != nil {
		user.Age = uint(*i.Age)
	}

	if s, ok := atmail.ValidateUser(user); !ok {
		return badRequest(s), nil
	}

	if err := s.UpdateUser(user); err != nil {
		return nil, err
	}

	o := updateUserOutload{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Age:      uint(user.Age),
	}

	return o, nil
}

type messageOutload struct {
	Message string `json:"message"`
}

func (o messageOutload) code() int {
	return http.StatusOK
}

func deleteUser(s atmail.Store, w http.ResponseWriter, r *http.Request) (outload, error) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		return nil, err
	}

	if err := s.DeleteUser(id); err != nil {
		if !errors.Is(err, atmail.ErrUserNone) {
			return nil, err
		}

		return badRequest("user does not exist!"), nil
	}

	return messageOutload{fmt.Sprintf("successfully deleted user %d!", id)}, nil
}
