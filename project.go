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

// ProjectWithIdentifier returns the Project with the given Identifier identifier.
// If no project found, returns nil.
func (r *Redmine) ProjectWithIdentifier(identifier string) (*Project, error) {

	projects, err := r.Projects(NameParameter(identifier))
	if err != nil {
		return nil, err
	}

	for i := range projects.Projects {
		if projects.Projects[i].Identifier == identifier {
			return &projects.Projects[i], nil
		}
	}

	return nil, nil
}
