package redmine

import (
	"encoding/json"
	"fmt"
	"time"
)

type Issue struct {
	ID          int      `json:"id"`
	Project     Project  `json:"project"`
	Tracker     Tracker  `json:"tracker"`
	Status      Status   `json:"status"`
	Priority    Priority `json:"priority"`
	Author      Author   `json:"author"`
	Category    Category `json:"category"`
	Subject     string   `json:"subject"`
	Description string   `json:"description"`
	StartDate   string   `json:"start_date"`
	DueDate     string   `json:"due_date"`
	DoneRation  int      `json:"done_ratio"`
	IsPrivate   bool     `json:"is_private"`
	//EstimatedHours
	//TotalEstimatedHours
	SpentHours      float64       `json:"spent_hours"`
	TotalSpentHours float64       `json:"total_spent_hours"`
	CreatedOn       time.Time     `json:"created_on"`
	UpdatedOn       time.Time     `json:"updated_on"`
	ClosedOn        time.Time     `json:"closed_on"`
	CustomFields    []CustomField `json:"custom_fields"`
}

func (i *Issue) String() string {
	return fmt.Sprintf("%s - %s #%d: %s", i.Project.Name, i.Tracker.Name, i.ID, i.Subject)
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
