package redmine_test

import (
	"fmt"
	"testing"

	"gorbe.io/go/redmine"
)

func ExampleIssues() {

	issues, err := redmine.NewPublic("https://www.redmine.org").Issues()
	if err != nil {
		// handle error
	}

	fmt.Printf("%#v", issues)
}

func ExampleIssues_withParameter() {

	issues, err := redmine.NewPublic("https://www.redmine.org").Issues(redmine.OffsetParameter(100), redmine.LimitParameter(100))
	if err != nil {
		// handle error
	}

	fmt.Printf("%#v", issues)
}

func TestIssues(t *testing.T) {

	issues, err := redmine.NewPublic("https://www.redmine.org").Issues()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	t.Logf("%s\n", issues.JSON())
}
