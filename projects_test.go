package redmine_test

import (
	"fmt"
	"testing"

	"gorbe.io/go/redmine"
)

func ExampleProjects() {

	v, err := redmine.NewPublic("https://www.redmine.org").Projects()
	if err != nil {
		// handle error
	}

	fmt.Printf("%v\n", v.Identifiers())
}

func TestProjects(t *testing.T) {

	v, err := redmine.NewPublic("https://www.redmine.org").Projects()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	t.Logf("%#v\n", v.Identifiers())
}
