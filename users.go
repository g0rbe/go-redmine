package redmine

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Users struct {
	Users      []User `json:"users,omitempty" yaml:"users,omitempty"`
	TotalCount int    `json:"total_count,omitempty" yaml:"total_count,omitempty"`
	Offset     int    `json:"offset,omitempty" yaml:"offset,omitempty"`
	Limit      int    `json:"limit,omitempty" yaml:"limit,omitempty"`
}

func (u *Users) String() string {

	var v string

	for i := range u.Users {
		v += fmt.Sprintf("%s\n", u.Users[i])
	}

	if len(u.Users) > 0 {
		v += "\n"
	}

	v += fmt.Sprintf("Total: %d\n", u.TotalCount)
	v += fmt.Sprintf("Offset: %d\n", u.Offset)
	v += fmt.Sprintf("Limit: %d", u.Limit)

	return v
}

// JSON encodes Users to JSON.
//
// If marshaling fails for any reason, this function panics.
func (u *Users) JSON() string {

	v, err := json.Marshal(u)
	if err != nil {
		panic(fmt.Errorf("failed to marshal Users to JSON: %w", err))
	}

	return string(v)
}

// YAML encodes Users to YAML.
//
// If marshaling fails for any reason, this function panics.
func (u *Users) YAML() string {

	v, err := yaml.Marshal(u)
	if err != nil {
		panic(fmt.Errorf("failed to marshal Users to YAML: %w", err))
	}

	if v[len(v)-1] == '\n' {
		v = v[:len(v)-1]
	}

	return string(v)
}

func (r *Redmine) Users(params ...Parameter) (*Users, error) {

	code, body, err := r.auth.Request("GET", "/users.json"+ParseParameters(params...), nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if code == 403 {
		return nil, ErrForbidden
	}
	if code != 200 {
		return nil, fmt.Errorf("(%d) %s", code, body)
	}

	v := new(Users)

	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return v, nil
}
