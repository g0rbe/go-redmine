package redmine

import (
	"encoding/json"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/yaml.v3"
)

type Users struct {
	Users      []User `json:"users,omitempty" yaml:"users,omitempty"`
	TotalCount int    `json:"total_count,omitempty" yaml:"total_count,omitempty"`
	Offset     int    `json:"offset,omitempty" yaml:"offset,omitempty"`
	Limit      int    `json:"limit,omitempty" yaml:"limit,omitempty"`
}

func (u *Users) ToTable() string {

	tw := table.NewWriter()

	tw.AppendHeader(table.Row{"#", "Login", "First Name", "Last Name", "Mail", "Last Login", "Admin", "Status"}, table.RowConfig{AutoMerge: true})

	for i := range u.Users {
		tw.AppendRow(table.Row{u.Users[i].ID, u.Users[i].Login, u.Users[i].FirstName, u.Users[i].LastName, u.Users[i].Mail, u.Users[i].LastLoginOn.Local(), u.Users[i].Admin, u.Users[i].Status.String()}, table.RowConfig{AutoMerge: true})
	}

	tw.AppendSeparator()

	tw.AppendFooter(table.Row{"", "", "", "", "", "", "Total", u.TotalCount})
	tw.AppendFooter(table.Row{"", "", "", "", "", "", "Offset", u.Offset})
	tw.AppendFooter(table.Row{"", "", "", "", "", "", "Limit", u.Limit})

	return tw.Render()
}

func (u *Users) ToJSON() string {

	v, err := json.Marshal(u)
	if err != nil {
		panic(fmt.Errorf("failed to marshal Users to JSON: %w", err))
	}

	return string(v)
}

func (u *Users) ToYAML() string {

	v, err := yaml.Marshal(u)
	if err != nil {
		panic(fmt.Errorf("failed to marshal Users to YAML: %w", err))
	}

	return string(v)
}
func (r *Redmine) Users(filter string) (*Users, error) {

	if len(filter) > 0 && filter[0] != '?' {
		filter = "?" + filter
	}

	code, body, err := r.auth.Request("GET", "/users.json"+filter, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if code == 403 {
		return nil, ErrForbidden
	}
	if code != 200 {
		return nil, fmt.Errorf("(%d) %s", code, body)
	}

	v := new(Users)

	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return v, nil
}
