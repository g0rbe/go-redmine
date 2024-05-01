package redmine

import (
	"encoding/json"
	"fmt"
)

type Projects struct {
	Projects   []Project `json:"projects,omitempty" yaml:"projects,omitempty"`
	TotalCount int       `json:"total_count,omitempty" yaml:"total_count,omitempty"`
	Offset     int       `json:"offset,omitempty" yaml:"offset,omitempty"`
	Limit      int       `json:"limit,omitempty" yaml:"limit,omitempty"`
}

// Identifiers returns a slice of projects identifier.
func (p *Projects) Identifiers() []string {

	v := make([]string, 0, len(p.Projects))

	for i := range p.Projects {
		v = append(v, p.Projects[i].Identifier)
	}

	return v
}

func (r *Redmine) Projects(params ...Parameter) (*Projects, error) {

	code, body, err := r.auth.Request("GET", "/projects.json"+ParseParameters(params...), nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if code == 403 {
		return nil, ErrForbidden
	}
	if code != 200 {
		return nil, fmt.Errorf("(%d) %s", code, body)
	}

	v := new(Projects)

	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return v, nil
}
