package redmine

import "time"

// TODO: Parent Project
type Project struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Identifier  string    `json:"identifier"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
	IsPublic    bool      `json:"is_public"`
	Homepage    string    `json:"homepage"`
}
