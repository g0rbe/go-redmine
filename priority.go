package redmine

import (
	"encoding/json"
	"fmt"
)

type Priority struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
}

type Priorities []Priority

func (p *Priorities) Names() []string {

	v := make([]string, 0, len(*p))

	for i := range *p {
		v = append(v, (*p)[i].Name)
	}

	return v
}

func (r *Redmine) Priorities(params ...Parameter) (Priorities, error) {

	code, body, err := r.auth.Request("GET", "/enumerations/issue_priorities.json"+ParseParameters(params...), nil)
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
		Priorities Priorities `json:"issue_priorities"`
	}{}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return v.Priorities, nil
}

// PriorityWithName returns the Priority with the given Name.
// If no Priority found with Name name, returns nil.
func (r *Redmine) PriorityWithName(name string) (*Priority, error) {

	ps, err := r.Priorities(NameParameter(name))
	if err != nil {
		return nil, err
	}

	for i := range ps {
		if ps[i].Name == name {
			return &ps[i], nil
		}
	}

	return nil, nil
}
