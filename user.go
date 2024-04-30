package redmine

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

type UserStatus int

const (
	UserActive     UserStatus = 0
	UserRegistered UserStatus = 1
	UserLocked     UserStatus = 2
)

// ParseUserStatus converts user status string v to UserStatus.
// Returns -1 if v is not a valid user status.
func ParseUserStatus(v string) UserStatus {
	switch v {
	case "active":
		return UserActive
	case "registered":
		return UserRegistered
	case "locked":
		return UserLocked
	default:
		return -1
	}
}

func (us UserStatus) String() string {
	switch us {
	case UserActive:
		return "active"
	case UserRegistered:
		return "registered"
	case UserLocked:
		return "locked"
	default:
		panic(fmt.Errorf("invalid value for UserStatus: %d", us))
	}
}

type MailNotification string

const (
	MailNotificationAll  MailNotification = "all"
	MailNotificationNone MailNotification = "none"
)

func (m MailNotification) String() string {
	switch m {
	case "all":
		return "For any event on all my projects"
	case "none":
		return "No events"
	default:
		return string(m)
	}
}

type User struct {
	ID              int           `json:"id,omitempty" yaml:"id,omitempty"`
	Login           string        `json:"login,omitempty" yaml:"login,omitempty"`
	Password        string        `json:"password,omitempty" yaml:"password,omitempty"`
	Admin           bool          `json:"admin,omitempty" yaml:"admin,omitempty"`
	FirstName       string        `json:"firstname,omitempty" yaml:"firstname,omitempty"`
	LastName        string        `json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Mail            string        `json:"mail,omitempty" yaml:"mail,omitempty"`
	CreatedOn       time.Time     `json:"created_on,omitempty" yaml:"created_on,omitempty"`
	UpdatedOn       time.Time     `json:"updated_on,omitempty" yaml:"updated_on,omitempty"`
	LastLoginOn     time.Time     `json:"last_login_on,omitempty" yaml:"last_login_on,omitempty"`
	PasswdChangedOn time.Time     `json:"passwd_changed_on,omitempty" yaml:"passwd_changed_on,omitempty"`
	TwoFAScheme     string        `json:"twofa_scheme,omitempty" yaml:"twofa_scheme,omitempty"`
	ApiKey          string        `json:"api_key,omitempty" yaml:"api_key,omitempty"`
	AvatarURL       string        `json:"avatar_url,omitempty" yaml:"avatar_url,omitempty"`
	Status          UserStatus    `json:"status,omitempty" yaml:"status,omitempty"`
	CustomField     []CustomField `json:"custom_fields,omitempty" yaml:"custom_fields,omitempty"`
	Groups          []Group       `json:"groups,omitempty" yaml:"groups,omitempty"`
	// Memberships

	// CreateUser fields
	AuthSourceId       int              `json:"auth_source_id,omitempty" yaml:"auth_source_id,omitempty"`
	MailNotification   MailNotification `json:"mail_notification,omitempty" yaml:"mail_notification,omitempty"`
	MustChangePassword bool             `json:"must_change_passwd,omitempty" yaml:"must_change_passwd,omitempty"`
	GeneratePassword   bool             `json:"generate_password,omitempty" yaml:"generate_password,omitempty"`
}

func (u User) String() string {
	return fmt.Sprintf("#%d %s %s (%s) %s %s", u.ID, u.FirstName, u.LastName, u.Login, u.Mail, u.LastLoginOn)
}

func (u User) ToYAML() []byte {

	v, err := yaml.Marshal(&u)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal user: %w", err))
	}

	return v
}

func (u User) ToJSON() []byte {

	v, err := json.Marshal(&u)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal user: %w", err))
	}

	return v
}

// If id is 0, then returns the current user
func (r *Redmine) User(id int) (User, error) {

	idStr := strconv.Itoa(id)

	if id == 0 {
		idStr = "current"
	}

	code, body, err := r.auth.Request("GET", fmt.Sprintf("/users/%s.json", idStr), nil)
	if err != nil {
		return User{}, fmt.Errorf("request failed: %w", err)
	}

	if code == 403 {
		return User{}, ErrForbidden
	}
	if code != 200 {
		return User{}, fmt.Errorf("(%d) %s", code, body)
	}

	v := struct {
		User User `json:"user"`
	}{}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return User{}, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return v.User, nil
}

// CreateUser creates a new user.
//
// Required fields: Login, Firstname, Lastname and Mail
//
// If notify is true, sends account information to the user.
// If user created, the underlying data of u will be replaced by the returned user.
func (r *Redmine) CreateUser(u *User, notify bool) error {

	if u == nil {
		return fmt.Errorf("user is nil")
	}
	if u.Login == "" {
		return ErrLoginEmpty
	}
	if u.FirstName == "" {
		return ErrLastnameEmpty
	}
	if u.LastName == "" {
		return ErrLastnameEmpty
	}
	if u.Mail == "" {
		return ErrMailEmpty
	}

	nu := struct {
		User            *User `json:"user"`
		SendInformation bool  `json:"send_information"`
	}{
		User:            u,
		SendInformation: notify,
	}

	code, body, err := r.auth.Request("POST", "/users.json", &nu)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if code == 422 {
		return fmt.Errorf("%w %s", ErrUnprocessableEntity, body)
	}
	if code != 201 {
		return fmt.Errorf("(%d) %s", code, body)
	}

	v := struct {
		User User `json:"user"`
	}{}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal body: %w", err)
	}

	*u = v.User

	return nil
}

// UpdateUser modify a user.
//
// Required fields: ID, Login, Firstname, Lastname and Mail
//
// If notify is true, sends account information to the user.
func (r *Redmine) UpdateUser(u *User, notify bool) error {

	if u == nil {
		return fmt.Errorf("user is nil")
	}
	if u.ID == 0 {
		return ErrIDEmpty
	}
	if u.Login == "" {
		return ErrLoginEmpty
	}
	if u.FirstName == "" {
		return ErrLastnameEmpty
	}
	if u.LastName == "" {
		return ErrLastnameEmpty
	}
	if u.Mail == "" {
		return ErrMailEmpty
	}

	nu := struct {
		User            *User `json:"user"`
		SendInformation bool  `json:"send_information"`
	}{
		User:            u,
		SendInformation: notify,
	}

	code, body, err := r.auth.Request("PUT", fmt.Sprintf("/users/%d.json", u.ID), &nu)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if code == 422 {
		return fmt.Errorf("%w %s", ErrUnprocessableEntity, body)
	}
	if code != 204 {
		return fmt.Errorf("(%d) %s", code, body)
	}

	return nil
}

// DeleteUser remves a user with the given id.
func (r *Redmine) DeleteUser(id int) error {

	code, body, err := r.auth.Request("DELETE", fmt.Sprintf("/users/%d.json", id), nil)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if code != 204 {
		return fmt.Errorf("(%d) %s", code, body)
	}

	return nil
}
