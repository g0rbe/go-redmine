package redmine

import (
	"encoding/json"
	"fmt"
)

type Tracker struct {
	ID                    int      `json:"id"`
	Name                  string   `json:"name"`
	DefaultStatus         Status   `json:"default_status"`
	Description           string   `json:"description"`
	EnabledStandardFields []string `json:"enabled_standard_fields"`
}

func (r *Redmine) Trackers() ([]Tracker, error) {

	code, body, err := r.auth.Request("GET", "/trackers.json", nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if code != 200 {
		return nil, fmt.Errorf("%d %s", code, body)
	}

	v := struct {
		Trackers []Tracker `json:"trackers"`
	}{}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return v.Trackers, nil
}
