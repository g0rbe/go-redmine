package redmine

import (
	"encoding/json"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/yaml.v3"
)

type Issues struct {
	Issues     []Issue `json:"issues,omitempty" yaml:"issues,omitempty"`
	TotalCount int     `json:"total_count,omitempty" yaml:"total_count,omitempty"`
	Offset     int     `json:"offset,omitempty" yaml:"offset,omitempty"`
	Limit      int     `json:"limit,omitempty" yaml:"limit,omitempty"`
}

func (i *Issues) ToTable() string {

	tw := table.NewWriter()

	tw.AppendHeader(table.Row{"#", "Project", "Tracker", "Status", "Priority", "Subject", "Start Date", "Due Date", "Total Spent Hours"}, table.RowConfig{AutoMerge: true})

	for n := range i.Issues {
		tw.AppendRow(
			table.Row{
				i.Issues[n].ID, i.Issues[n].Project, i.Issues[n].Tracker.Name, i.Issues[n].Status, i.Issues[n].Priority,
				i.Issues[n].Subject, i.Issues[n].StartDate, i.Issues[n].DueDate, i.Issues[n].TotalSpentHours,
			},
			table.RowConfig{AutoMerge: true})
	}

	tw.AppendFooter(table.Row{"", "", "", "", "", "", "Total", i.TotalCount})
	tw.AppendFooter(table.Row{"", "", "", "", "", "", "Offset", i.Offset})
	tw.AppendFooter(table.Row{"", "", "", "", "", "", "Limit", i.Limit})

	return tw.Render()
}

func (i *Issues) ToJSON() (string, error) {

	v, err := json.Marshal(i)
	if err != nil {
		return "", (fmt.Errorf("failed to marshal Issues to JSON: %w", err))
	}

	return string(v), nil
}

func (i *Issues) MustToJSON() string {

	v, err := i.ToJSON()
	if err != nil {
		panic(fmt.Errorf("failed to marshal Issues to JSON: %w", err))
	}

	return v
}

func (i *Issues) ToYAML() (string, error) {

	v, err := yaml.Marshal(i)
	if err != nil {
		return "", (fmt.Errorf("failed to marshal Issues to YAML: %w", err))
	}

	return string(v), nil
}

func (i *Issues) MustToYAML() string {

	v, err := i.ToYAML()
	if err != nil {
		panic(fmt.Errorf("failed to marshal Issues to YAML: %w", err))
	}

	return v
}

func (r *Redmine) Issues(filter string, limit int, offset int) (*Issues, error) {

	if limit == 0 {
		limit = 25
	}

	if filter != "" {
		filter += "&"
	}
	code, body, err := r.auth.Request("GET", fmt.Sprintf("/issues.json?%slimit=%d&offset=%d", filter, limit, offset), nil)
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
