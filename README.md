# Atmail

This is a simple CRUD API server for users.

## Tasks
1. User Management API:
    - Endpoints:
        - [x] `POST /users`
        - [x] `GET /users/{id}`
        - [x] `PUT /users/{id}`
        - [x] `DELETE /users/{id}`
    - User Data:
        - [x] `id`
        - [x] `username`
        - [x] `email`
        - [x] `age`
    - [x] Validation
- [x] 2. Authentication
3. Database
    - [x] Use MySQL
    - [x] Create `users` table
    - [x] Set up a connection pool
4. Error Handling
    - [x] Handle errors gracefully
    - [x] Return JSON responses with error messages
- [x] 5. Frameworks
- [x] 6. Unit tests
- [x] 7. RBAC
- [x] 8. API Documentation (OpenAPI 3)

## Setup

Run the following commands to get setup:

```plaintext
$ git clone git@github.com:luiscvega/atmail.git
$ cd atmail
$ go mod tidy
$ mysql -u <user> -p -e 'CREATE DATABASE IF NOT EXISTS atmail'
$ mysql -u <user> -p atmail < setup.sql
```

### Running the server

```plaintext
$ MYSQL_URL=<mysql-url> PORT=8080 go run cmd/atmail/main.go
```

You can then run a curl command to create a user:

```plaintext
$ curl localhost:8080/users -d '{ "username": "johndoe", "email": "john@doe.com", "age": 123 }' \
                            -u bob:pass2345
{
	"id": 1,
	"username": "johndoe",
	"email": "john@doe.com",
	"age": 123
}

```

### Running tests

```plaintext
$ MYSQL_URL=<mysql-url> go test ./... 
```

## Structure

### `cmd/atmail/`

This is the main program that will run the server.

### `server/`

This is the server layer that contains everything related to the server such as handlers and the roles. I have also designed a simple `handler` that implements the `http.Handler` interface for better control over the application.

### `atmail.go`

This is the main business logic when interacting with the API server. Since this is a simple CRUD app, most of the functions here wrappers for database calls.

The key part here is that we have created a `Store` interface that can be subsituted with any database implementation (the default is MySQL). Of course, this allows us to mock calls for easier and faster tests

### `api.yaml`

This is the Open API 3 specification for the API server. This makes is easier for others to implement clients for this server.

### `api/`

This is the generated API client in Go from the `api.yaml` file. This also contains a test to ensure that the API is designed correctly.

## Roles

Roles are designed by using bitwise operations. Here is an example of a handler that specifies the permissions on it.

```go
mux.Handle("POST /users", toHandler(createUser,
	roles.Bluey|roles.Chilli|roles.Bandit,
))
```

As you can see, you can fine-tune the roles that are allowed to access the specific handler by using the `|` bitwise operator.

You can refer to this table below for the roles included in the SQL dump:

| Endpoint             | Bingo  | Bluey  | Chilli | Bandit |
|----------------------|--------|--------|--------|--------|
| `GET /users/{id}`    | ✔      | ✔      | ✔      | ✔      |
| `POST /users`        | ❌     | ✔      | ✔      | ✔      |
| `PUT /users/{id}`    | ❌     | ❌     | ✔      | ✔      |
| `DELETE /users/{id}` | ❌     | ❌     | ❌     | ✔      |

## Admins

Admins are included in the `setup.sql` file and created when the project is setup. 

This is a table of admins with varying roles assigned to them:

| User  | Password | Role   |
|-------|----------|--------|
| alice | pass1234 | Bingo  |
| bob   | pass2345 | Bluey  |
| craig | pass3456 | Chilli |
| dan   | pass4567 | Bandit |


## API Documentation

This API provides endpoints to manage users, including creating, retrieving, updating, and deleting users. The API follows OpenAPI 3.0.2 specifications and supports basic authentication for security.

### Endpoints

#### 1. **Create a new user**
- **URL:** `/users`
- **Method:** `POST`
- **Summary:** Creates a new user with provided `username`, `email`, and `age`.
- **Request Body:**
  - `Content-Type`: `application/json`
  - **Required Properties:**
    - `username` (string)
    - `email` (string)
    - `age` (integer, int64)
- **Responses:**
  - **201 Created**: Returns the created user object.
      ##### Example Request:
      ```json
      {
        "username": "john_doe",
        "email": "john.doe@example.com",
        "age": 30
      }
      ```

      ##### Example Response:
      ```json
      {
        "id": 1,
        "username": "john_doe",
        "email": "john.doe@example.com",
        "age": 30
      }
      
      ```
  - **400 Bad Request**: Request failed due to invalid user input.
      ##### Example Request:
      ```json
      {
        "username": "",
        "email": "john.doe.updated@example.com",
        "age": 31
      }
      ```

      ##### Example Response:
      ```json
      {
        "error": "username cannot be blank!"
      }
      ```
  - **401 Unauthorized**: Authentication failed or user does not have permission to perform this action.
  - **500 Internal Server Error**: A general server error occurred.

#### 2. **Get a user**
- **URL:** `/users/{id}`
- **Method:** `GET`
- **Summary:** Retrieves the details of a user by their `id`.
- **Parameters:**
  - `id` (path parameter): The ID of the user to retrieve.
- **Responses:**
  - **200 OK**: Returns the user object for the requested `id`.

    ##### Example Response:
    ```json
    {
      "id": 1,
      "username": "john_doe",
      "email": "john.doe@example.com",
      "age": 30
    }
    ```

  - **400 Bad Request**: Request failed because user ID does not exist
    ##### Example Response:
    ```json
    {
      "error": "user does not exist!"
    }
    ```
  - **401 Unauthorized**: Authentication failed or user does not have permission to perform this action.
  - **500 Internal Server Error**: A general server error occurred.


#### 3. **Update a user**
- **URL:** `/users/{id}`
- **Method:** `PUT`
- **Summary:** Updates the details of a user identified by their `id`.
- **Parameters:**
  - `id` (path parameter): The ID of the user to update.
- **Request Body:**
  - `Content-Type`: `application/json`
  - **Required Properties:**
    - `username` (string)
    - `email` (string)
    - `age` (integer, int64)
- **Responses:**
  - **200 OK**: Returns the updated user object.
    ##### Example Request:
    ```json
    {
      "username": "john_doe_updated",
      "email": "john.doe.updated@example.com",
      "age": 31
    }
    ```
    
    ##### Example Response:
    ```json
    {
      "id": 1,
      "username": "john_doe_updated",
      "email": "john.doe.updated@example.com",
      "age": 31
    }
    ```

  - **400 Bad Request**: Request failed due to invalid user input.
    ##### Example Request:
    ```json
    {
      "username": "",
      "email": "john.doe.updated@example.com",
      "age": 31
    }
    ```
    
    ##### Example Response:
    ```json
    {
      "error": "username cannot be blank!"
    }
    ```
  - **401 Unauthorized**: Authentication failed or user does not have permission to perform this action.
  - **500 Internal Server Error**: A general server error occurred.

#### 4. **Delete a user**
- **URL:** `/users/{id}`
- **Method:** `DELETE`
- **Summary:** Deletes a user by their `id`.
- **Parameters:**
  - `id` (path parameter): The ID of the user to delete.
- **Responses:**
  - **200 OK**: Returns a message confirming the deletion.
    ##### Example Response:
    ```json
    {
      "message": "successfully deleted user 1234!"
    }
    ```

  - **400 Bad Request**: Request failed because user ID does not exist
    ##### Example Response:
    ```json
    {
      "error": "user does not exist!"
    }
    ```
  - **401 Unauthorized**: Authentication failed or user does not have permission to perform this action.
  - **500 Internal Server Error**: A general server error occurred.

### Components

#### User Schema

The following schema is used for representing a user in the system.

##### Properties:
- `id` (integer, int64): The unique identifier for the user.
- `username` (string): The username of the user.
- `email` (string): The email address of the user.
- `age` (integer, int64): The age of the user.

##### Example User Object:
```json
{
  "id": 1,
  "username": "john_doe",
  "email": "john.doe@example.com",
  "age": 30
}
```

#### Response Objects

##### Bad Request Response:
This response is returned when the request does not have valid request body.

```json
{
  "error": "username cannot be blank!"
}
```

##### Unauthorized Response:
This response is returned when the request does not have valid authentication or the user does not have proper permissions.

```json
{
  "error": "unauthorized!"
}
```

##### Internal Server Error Response:
This response is returned when there is an error processing the request on the server.

```json
{
  "error": "internal server error!"
}
```

### Security

This API uses Basic Authentication for all endpoints.

- **Security Scheme:** Basic HTTP Authentication

#### Example Header:
```plaintext
Authorization: Basic <base64_encoded_credentials>
```

Where `<base64_encoded_credentials>` is the base64 encoding of `username:password`.

### Errors

The API uses standard HTTP status codes to indicate the outcome of API requests:

- **2xx**: Success (e.g., 201 Created, 200 OK)
- **4xx**: Client error (e.g., 401 Unauthorized)
- **5xx**: Server error (e.g., 500 Internal Server Error)
