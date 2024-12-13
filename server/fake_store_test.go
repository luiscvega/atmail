package server

import (
	"atmail"
	"atmail/server/roles"
)

type fakeStore struct {
	newUserId             int64
	existingUser          atmail.User
	existingAdminUser     string
	existingAdminPassword string
	existingAdminRole     roles.Role
}

func (s fakeStore) CheckUser(username string, email string) (bool, error) {
	return s.existingUser.Username == username || s.existingUser.Email == email, nil
}

func (s fakeStore) CreateUser(atmail.User) (int64, error) {
	return s.newUserId, nil
}

func (s fakeStore) GetUser(id int64) (atmail.User, error) {
	if id != s.existingUser.Id {
		return atmail.User{}, atmail.ErrUserNone
	}

	return s.existingUser, nil
}

func (s fakeStore) UpdateUser(atmail.User) error {
	return nil
}

func (s fakeStore) DeleteUser(id int64) error {
	return nil
}

func (s fakeStore) GetRole(user string, password string) (roles.Role, error) {
	if s.existingAdminUser != user && s.existingAdminPassword != password {
		return 0, atmail.ErrAdminNone
	}

	return s.existingAdminRole, nil
}
