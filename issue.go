package redmine

import (
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
)

type Issue struct {
	ID                  int           `json:"id,omitempty" yaml:"id,omitempty"`
	Project             *Project      `json:"project,omitempty" yaml:"project,omitempty"`
	Tracker             *Tracker      `json:"tracker,omitempty" yaml:"tracker,omitempty"`
	Status              *IssueStatus  `json:"status,omitempty" yaml:"status,omitempty"`
	Priority            *Priority     `json:"priority,omitempty" yaml:"priority,omitempty"`
	Author              *Author       `json:"author,omitempty" yaml:"author,omitempty"`
	AssignedTo          *User         `json:"assigned_to,omitempty" yaml:"assigned_to,omitempty"`
	Category            *Category     `json:"category,omitempty" yaml:"category,omitempty"`
	Subject             string        `json:"subject,omitempty" yaml:"subject,omitempty"`
	Description         string        `json:"description,omitempty" yaml:"description,omitempty"`
	StartDate           string        `json:"start_date,omitempty" yaml:"start_date,omitempty"`
	DueDate             string        `json:"due_date,omitempty" yaml:"due_date,omitempty"`
	DoneRation          int           `json:"done_ratio,omitempty" yaml:"done_ratio,omitempty"`
	IsPrivate           bool          `json:"is_private,omitempty" yaml:"is_private,omitempty"`
	SpentHours          float64       `json:"spent_hours,omitempty" yaml:"spent_hours,omitempty"`
	TotalSpentHours     float64       `json:"total_spent_hours,omitempty" yaml:"total_spent_hours,omitempty"`
	CreatedOn           *time.Time    `json:"created_on,omitempty" yaml:"created_on,omitempty"`
	UpdatedOn           *time.Time    `json:"updated_on,omitempty" yaml:"updated_on,omitempty"`
	ClosedOn            *time.Time    `json:"closed_on,omitempty" yaml:",omitempty"`
	CustomFields        []CustomField `json:"custom_fields,omitempty" yaml:"custom_fields,omitempty"`
	EstimatedHours      float64       `json:"estimated_hours,omitempty" yaml:"estimated_hours,omitempty"`
	TotalEstimatedHours float64       `json:"total_estimated_hours,omitempty" yaml:"total_estimated_hours,omitempty"`
}

func (i *Issue) String() string {
	return fmt.Sprintf("%s %s %s #%d: %s", i.Status.Name, i.Project.Name, i.Tracker.Name, i.ID, i.Subject)
}

func (i *Issue) YAML() string {

	v, err := yaml.Marshal(i)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal Issue: %w", err))
	}

	if v[len(v)-1] == '\n' {
		v = v[:len(v)-1]
	}

	return string(v)
}

func (i *Issue) JSON() string {

	v, err := json.Marshal(i)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal Issue: %w", err))
	}

	return string(v)
}

func (r *Redmine) Issue(id int) (*Issue, error) {

	code, body, err := r.auth.Request("GET", fmt.Sprintf("/issues/%d.json", id), nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if code != 200 {
		return nil, fmt.Errorf("(%d) %s", code, body)
	}

	v := struct {
		Issue Issue `json:"issue"`
	}{}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return &v.Issue, nil
}

// CreateIssue creates a new Issue.
//
// If Issue is created, the underlying data of iss will be replaced by the returned Issue.
func (r *Redmine) CreateIssue(i *Issue) error {

	ni := struct {
		Issue *Issue `json:"issue"`
	}{
		Issue: i,
	}

	code, body, err := r.auth.Request("POST", "/issues.json", &ni)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if code == 422 {
		return fmt.Errorf("%w %s", ErrUnprocessableEntity, body)
	}
	if code != 201 {
		return fmt.Errorf("(%d) %s", code, body)
	}

	err = json.Unmarshal(body, &ni)
	if err != nil {
		return fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return nil
}
