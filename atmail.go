package atmail

import (
	"database/sql"
	"errors"
	"net/mail"

	"atmail/server/roles"
)

var ErrUserNone = errors.New("error user none")
var ErrAdminNone = errors.New("error admin none")

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return store{db}
}

type Store interface {
	CheckUser(string, string) (bool, error)
	CreateUser(User) (int64, error)
	GetUser(int64) (User, error)
	UpdateUser(User) error
	DeleteUser(int64) error

	GetRole(string, string) (roles.Role, error)
}

type User struct {
	Id       int64
	Username string
	Email    string
	Age      uint
}

func ValidateUser(user User) (string, bool) {
	if user.Username == "" {
		return "username cannot be blank!", false
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return "email is invalid!", false
	}

	if user.Age <= 0 {
		return "age is invalid!", false
	}

	return "", true
}

func (s store) GetUser(id int64) (User, error) {
	user := User{}

	if err := s.db.QueryRow("SELECT id, username, email, age FROM users WHERE id = ?", id).Scan(&user.Id, &user.Username, &user.Email, &user.Age); err != nil {
		if err != sql.ErrNoRows {
			return User{}, err
		}

		return User{}, ErrUserNone
	}

	return user, nil
}

func (s store) CheckUser(username string, password string) (bool, error) {
	var exists bool

	if err := s.db.QueryRow("SELECT count(*) != 0 FROM users WHERE username = ? OR email = ?", username, password).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (s store) CreateUser(user User) (int64, error) {
	result, err := s.db.Exec("INSERT INTO users (username, email, age) VALUES (?, ?, ?)", user.Username, user.Email, user.Age)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s store) UpdateUser(user User) error {
	if _, err := s.db.Exec("UPDATE users SET username = ?, email = ?, age = ? WHERE id = ?", user.Username, user.Email, user.Age, user.Id); err != nil {
		return err
	}

	return nil
}

func (s store) DeleteUser(id int64) error {
	result, err := s.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count != 1 {
		return ErrUserNone
	}

	return nil
}

func (s store) GetRole(user string, password string) (roles.Role, error) {
	var role roles.Role

	if err := s.db.QueryRow("SELECT role FROM admins WHERE user = ? AND password = ?", user, password).Scan(&role); err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}

		return 0, ErrAdminNone
	}

	return role, nil
}
