package redmine

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Issues struct {
	Issues     []Issue `json:"issues,omitempty" yaml:"issues,omitempty"`
	TotalCount int     `json:"total_count,omitempty" yaml:"total_count,omitempty"`
	Offset     int     `json:"offset,omitempty" yaml:"offset,omitempty"`
	Limit      int     `json:"limit,omitempty" yaml:"limit,omitempty"`
}

// JSON encodes Issues to JSON.
//
// If marshaling fails for any reason, this function panics.
func (i *Issues) JSON() string {

	v, err := json.Marshal(i)
	if err != nil {
		panic(fmt.Errorf("failed to marshal Issues to JSON: %w", err))
	}

	return string(v)
}

// YAML encodes Issues to YAML.
//
// If marshaling fails for any reason, this function panics.
func (i *Issues) YAML() string {

	v, err := yaml.Marshal(i)
	if err != nil {
		panic(fmt.Errorf("failed to marshal Issues to YAML: %w", err))
	}

	return string(v)
}

func (r *Redmine) Issues(params ...Parameter) (*Issues, error) {

	code, body, err := r.auth.Request("GET", "/issues.json"+ParseParameters(params...), nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if code != 200 {
		return nil, fmt.Errorf("(%d) %s", code, body)
	}

	//fmt.Printf("%s\n", body)
	v := new(Issues)

	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return v, nil
}
