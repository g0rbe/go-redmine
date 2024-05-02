package redmine

import (
	"encoding/json"
	"fmt"
)

type IssueStatus struct {
	ID          int    `json:"id,omitempty" yaml:"id,omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	IsClosed    bool   `json:"is_closed,omitempty" yaml:"is_closed,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type IssueStatuses []IssueStatus

func (is *IssueStatuses) Names() []string {

	v := make([]string, 0, len(*is))

	for i := range *is {
		v = append(v, (*is)[i].Name)
	}

	return v
}

func (r *Redmine) IssueStatuses(params ...Parameter) (IssueStatuses, error) {

	code, body, err := r.auth.Request("GET", "/issue_statuses.json"+ParseParameters(params...), nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if code == 403 {
		return nil, ErrForbidden
	}
	if code != 200 {
		return nil, fmt.Errorf("(%d) %s", code, body)
	}

	v := struct {
		IssueStatuses IssueStatuses `json:"issue_statuses"`
	}{}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return v.IssueStatuses, nil
}

// IssueStatusWithName returns the IssueStatus with the given Name.
// If no IssueStatus found with Name name, returns nil.
func (r *Redmine) IssueStatusWithName(name string) (*IssueStatus, error) {

	is, err := r.IssueStatuses(NameParameter(name))
	if err != nil {
		return nil, err
	}

	for i := range is {
		if is[i].Name == name {
			return &is[i], nil
		}
	}

	return nil, nil
}
