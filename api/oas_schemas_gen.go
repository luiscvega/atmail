// Code generated by ogen, DO NOT EDIT.

package api

// Ref: #/components/responses/badRequest
type BadRequest struct{}

func (*BadRequest) createUserRes() {}
func (*BadRequest) deleteUserRes() {}
func (*BadRequest) getUserRes()    {}
func (*BadRequest) updateUserRes() {}

type BasicAuth struct {
	Username string
	Password string
}

// GetUsername returns the value of Username.
func (s *BasicAuth) GetUsername() string {
	return s.Username
}

// GetPassword returns the value of Password.
func (s *BasicAuth) GetPassword() string {
	return s.Password
}

// SetUsername sets the value of Username.
func (s *BasicAuth) SetUsername(val string) {
	s.Username = val
}

// SetPassword sets the value of Password.
func (s *BasicAuth) SetPassword(val string) {
	s.Password = val
}

type CreateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int64  `json:"age"`
}

// GetUsername returns the value of Username.
func (s *CreateUserReq) GetUsername() string {
	return s.Username
}

// GetEmail returns the value of Email.
func (s *CreateUserReq) GetEmail() string {
	return s.Email
}

// GetAge returns the value of Age.
func (s *CreateUserReq) GetAge() int64 {
	return s.Age
}

// SetUsername sets the value of Username.
func (s *CreateUserReq) SetUsername(val string) {
	s.Username = val
}

// SetEmail sets the value of Email.
func (s *CreateUserReq) SetEmail(val string) {
	s.Email = val
}

// SetAge sets the value of Age.
func (s *CreateUserReq) SetAge(val int64) {
	s.Age = val
}

type DeleteUserOK struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *DeleteUserOK) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *DeleteUserOK) SetMessage(val string) {
	s.Message = val
}

func (*DeleteUserOK) deleteUserRes() {}

// Ref: #/components/responses/internalServerError
type InternalServerError struct{}

func (*InternalServerError) createUserRes() {}
func (*InternalServerError) deleteUserRes() {}
func (*InternalServerError) getUserRes()    {}
func (*InternalServerError) updateUserRes() {}

// NewOptInt64 returns new OptInt64 with value set to v.
func NewOptInt64(v int64) OptInt64 {
	return OptInt64{
		Value: v,
		Set:   true,
	}
}

// OptInt64 is optional int64.
type OptInt64 struct {
	Value int64
	Set   bool
}

// IsSet returns true if OptInt64 was set.
func (o OptInt64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt64) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt64) Get() (v int64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt64) Or(d int64) int64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/responses/unauthorized
type Unauthorized struct{}

func (*Unauthorized) createUserRes() {}
func (*Unauthorized) deleteUserRes() {}
func (*Unauthorized) getUserRes()    {}
func (*Unauthorized) updateUserRes() {}

type UpdateUserReq struct {
	Username OptString `json:"username"`
	Email    OptString `json:"email"`
	Age      OptInt64  `json:"age"`
}

// GetUsername returns the value of Username.
func (s *UpdateUserReq) GetUsername() OptString {
	return s.Username
}

// GetEmail returns the value of Email.
func (s *UpdateUserReq) GetEmail() OptString {
	return s.Email
}

// GetAge returns the value of Age.
func (s *UpdateUserReq) GetAge() OptInt64 {
	return s.Age
}

// SetUsername sets the value of Username.
func (s *UpdateUserReq) SetUsername(val OptString) {
	s.Username = val
}

// SetEmail sets the value of Email.
func (s *UpdateUserReq) SetEmail(val OptString) {
	s.Email = val
}

// SetAge sets the value of Age.
func (s *UpdateUserReq) SetAge(val OptInt64) {
	s.Age = val
}

// Ref: #/components/schemas/user
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int64  `json:"age"`
}

// GetID returns the value of ID.
func (s *User) GetID() int64 {
	return s.ID
}

// GetUsername returns the value of Username.
func (s *User) GetUsername() string {
	return s.Username
}

// GetEmail returns the value of Email.
func (s *User) GetEmail() string {
	return s.Email
}

// GetAge returns the value of Age.
func (s *User) GetAge() int64 {
	return s.Age
}

// SetID sets the value of ID.
func (s *User) SetID(val int64) {
	s.ID = val
}

// SetUsername sets the value of Username.
func (s *User) SetUsername(val string) {
	s.Username = val
}

// SetEmail sets the value of Email.
func (s *User) SetEmail(val string) {
	s.Email = val
}

// SetAge sets the value of Age.
func (s *User) SetAge(val int64) {
	s.Age = val
}

func (*User) createUserRes() {}
func (*User) getUserRes()    {}
func (*User) updateUserRes() {}
