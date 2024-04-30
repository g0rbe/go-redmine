package redmine

import (
	"fmt"
	"strconv"
)

type Parameter struct {
	Field string
	Value string
}

func (p Parameter) String() string {
	return fmt.Sprintf("%s=%s", p.Field, p.Value)
}

func ParseParameters(params ...Parameter) string {

	var v string

	for i := range params {
		if i == 0 {
			v += "?"
		} else {
			v += "&"
		}

		v += params[i].String()
	}

	return v
}

func OffsetParameter(v int) Parameter {
	return Parameter{Field: "offset", Value: strconv.Itoa(v)}
}

func LimitParameter(v int) Parameter {
	return Parameter{Field: "limit", Value: strconv.Itoa(v)}
}

func IncludeParameter(v string) Parameter {
	return Parameter{Field: "include", Value: v}
}

func IssueIDFilter(v int) Parameter {
	return Parameter{Field: "issue_id", Value: strconv.Itoa(v)}
}

func ProjectIDFilter(v int) Parameter {
	return Parameter{Field: "project_id", Value: strconv.Itoa(v)}
}

func StatusIDFilter(v int) Parameter {
	return Parameter{Field: "status_id", Value: strconv.Itoa(v)}
}

func StatusFilter(v int) Parameter {
	return Parameter{Field: "status", Value: strconv.Itoa(v)}
}

func NameParameter(v string) Parameter {
	return Parameter{Field: "name", Value: v}
}

func GroupIDFilter(v int) Parameter {
	return Parameter{Field: "group_id", Value: strconv.Itoa(v)}
}
