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

type User struct {
	ID              int           `json:"id,omitempty" yaml:"id,omitempty"`
	Login           string        `json:"login,omitempty" yaml:"login,omitempty"`
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

	var idStr string
	if id == 0 {
		idStr = "current"
	} else {
		strconv.Itoa(id)
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

	fmt.Printf("%s\n", body)

	v := struct {
		User User `json:"user"`
	}{}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return User{}, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return v.User, nil
}
