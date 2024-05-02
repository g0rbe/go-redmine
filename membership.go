package redmine

type Membership struct {
	Id      int      `json:"id,omitempty" yaml:"id,omitempty"`
	Project *Project `json:"project,omitempty" yaml:"project,omitempty"`
	User    *User    `json:"user,omitempty" yaml:"user,omitempty"`
	Roles   []Role   `json:"roles,omitempty" yaml:"roles,omitempty"`
}
