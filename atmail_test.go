package atmail_test

import (
	"atmail"
	"database/sql"
	"os"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestAtmail(t *testing.T) {
	// Step 1: Setup store
	databaseUrl := os.Getenv("MYSQL_URL")

	if databaseUrl == "" {
		t.Fatal("MYSQL_URL cannot be blank!")
	}

	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := db.Exec("DELETE FROM users"); err != nil {
		t.Fatal(err)
	}

	s := atmail.NewStore(db)

	// Step 2: Create user
	user := atmail.User{
		Username: "johndoe",
		Email:    "john@doe.com",
		Age:      55,
	}

	id, err := s.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	user.Id = id

	// Step 2: Get user
	gotUser, err := s.GetUser(id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(user, gotUser) {
		t.Errorf("want %v; got %v", user, gotUser)
	}

	// Step 3: Update user
	if err := s.UpdateUser(atmail.User{
		Username: "jane",
		Email:    "jane@doe.com",
		Age:      321,
	}); err != nil {
		t.Error(err)
	}

	// Step 4: Delete user
	if err := s.DeleteUser(id); err != nil {
		t.Error(err)
	}
}
