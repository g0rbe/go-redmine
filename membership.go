package redmine

type Membership struct {
	Project Project `json:"project"`
	Roles   []Role  `json:"roles"`
}
