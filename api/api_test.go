package api_test

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest -package api -target ../api --clean ../api.yaml

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"atmail"
	"atmail/api"
	"atmail/server"

	_ "github.com/go-sql-driver/mysql"
)

type basicAuth struct {
	username string
	password string
}

func (b basicAuth) BasicAuth(_ context.Context, _ api.OperationName) (api.BasicAuth, error) {
	return api.BasicAuth{b.username, b.password}, nil
}

func TestApi(t *testing.T) {
	// setup store
	databaseUrl := os.Getenv("MYSQL_URL")

	if databaseUrl == "" {
		log.Fatal("MYSQL_URL cannot be blank!")
	}

	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := db.Exec("DELETE FROM users"); err != nil {
		t.Fatal(err)
	}

	store := atmail.NewStore(db)

	// setup server
	server := httptest.NewServer(server.New(store))
	defer server.Close()

	// Step 1: Prepare api
	c, err := api.NewClient(server.URL, basicAuth{"dan", "pass4567"})
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	// Step 2: Create user
	createUserRes, err := c.CreateUser(ctx, &api.CreateUserReq{
		Username: "validusername",
		Email:    "valid@email.com",
		Age:      123,
	})
	if err != nil {
		t.Fatal(err)
	}

	user, ok := createUserRes.(*api.User)
	if !ok {
		t.Error("response is not a user!")
	}

	// Step 2: Get user
	getUserRes, err := c.GetUser(ctx, api.GetUserParams{ID: user.ID})
	if err != nil {
		t.Error(err)
	}

	gotUser, ok := getUserRes.(*api.User)
	if !ok {
		t.Error("response is not a user!")
	} else if !reflect.DeepEqual(user, gotUser) {
		t.Error("user is not the same!")
	}

	// Step 3: Update user
	updateUserRes, err := c.UpdateUser(ctx, &api.UpdateUserReq{
		Username: api.NewOptString("validusername2"),
		Email:    api.NewOptString("valid2@email.com"),
		Age:      api.NewOptInt64(321),
	}, api.UpdateUserParams{ID: user.ID})
	if err != nil {
		t.Error(err)
	}

	gotUser, ok = updateUserRes.(*api.User)
	if !ok {
		t.Error("response is not a user!")
	}

	if gotUser.Username != "validusername2" {
		t.Error("username is wrong")
	}

	if gotUser.Email != "valid2@email.com" {
		t.Error("email is wrong")
	}

	if gotUser.Age != 321 {
		t.Error("age is wrong")
	}

	// Step 4: Delete user
	deleteUserRes, err := c.DeleteUser(ctx, api.DeleteUserParams{ID: user.ID})
	if err != nil {
		t.Error(err)
	}

	deleteUserOk, ok := deleteUserRes.(*api.DeleteUserOK)
	if !ok {
		t.Error("response is not deleteUserOk")
	}

	if deleteUserOk.Message != fmt.Sprintf("successfully deleted user %d!", user.ID) {
		t.Error("delete message is wrong!")
	}
}
