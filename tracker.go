package redmine

import (
	"encoding/json"
	"fmt"
)

type Tracker struct {
	ID                    int         `json:"id,omitempty" yaml:"id,omitempty"`
	Name                  string      `json:"name,omitempty" yaml:"name,omitempty"`
	DefaultStatus         IssueStatus `json:"default_status,omitempty" yaml:"default_status,omitempty"`
	Description           string      `json:"description,omitempty" yaml:"description,omitempty"`
	EnabledStandardFields []string    `json:"enabled_standard_fields,omitempty" yaml:"enabled_standard_fields,omitempty"`
}

type Trackers []Tracker

func (t *Trackers) Names() []string {

	v := make([]string, 0, len(*t))

	for i := range *t {
		v = append(v, (*t)[i].Name)
	}

	return v
}

func (r *Redmine) Trackers(params ...Parameter) (Trackers, error) {

	code, body, err := r.auth.Request("GET", "/trackers.json"+ParseParameters(params...), nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if code != 200 {
		return nil, fmt.Errorf("%d %s", code, body)
	}

	v := struct {
		Trackers Trackers `json:"trackers"`
	}{}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return v.Trackers, nil
}

// TrackerWithName returns the Tracker with the given Name.
// If no Tracker found with Name name, returns nil.
func (r *Redmine) TrackerWithName(name string) (*Tracker, error) {

	trks, err := r.Trackers(NameParameter(name))
	if err != nil {
		return nil, err
	}

	for i := range trks {
		if trks[i].Name == name {
			return &trks[i], nil
		}
	}

	return nil, nil
}
