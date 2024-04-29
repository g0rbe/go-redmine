package redmine_test

import (
	"testing"

	"gorbe.io/go/redmine"
)

func TestIssues(t *testing.T) {

	issues, err := redmine.NewPublic("https://www.redmine.org").Issues("", 0, 0)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	t.Logf("%s\n", issues.MustToJSON())
}
