package redmine

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var ErrProjectNotFound = errors.New("project not found")

// TODO: Parent Project
type Project struct {
	ID          int       `json:"id,omitempty" yaml:"id,omitempty"`
	Name        string    `json:"name,omitempty" yaml:"name,omitempty"`
	Identifier  string    `json:"identifier,omitempty" yaml:"identifier,omitempty"`
	Description string    `json:"description,omitempty" yaml:"description,omitempty"`
	CreatedOn   time.Time `json:"created_on,omitempty" yaml:"created_on,omitempty"`
	UpdatedOn   time.Time `json:"updated_on,omitempty" yaml:"updated_on,omitempty"`
	IsPublic    bool      `json:"is_public,omitempty" yaml:"is_public,omitempty"`
	Homepage    string    `json:"homepage,omitempty" yaml:"homepage,omitempty"`
}

// ProjectWithIdentifier returns the Project with the given Identifier identifier.
//
// If no project found with the given identifier returns ErrProjectNotFound.
func (r *Redmine) ProjectWithIdentifier(identifier string) (*Project, error) {

	code, body, err := r.auth.Request("GET", fmt.Sprintf("/projects/%s.json", identifier), nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if code == 403 {
		return nil, ErrForbidden
	}
	if code == 404 {
		return nil, fmt.Errorf("not found")
	}
	if code != 200 {
		return nil, fmt.Errorf("(%d) %s", code, body)
	}

	v := struct {
		Project *Project `json:"project"`
	}{}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return v.Project, nil
}
