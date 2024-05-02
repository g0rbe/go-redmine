package redmine_test

import (
	"os"
	"testing"

	"gorbe.io/go/redmine"
)

func TestTrackers(t *testing.T) {

	v, err := redmine.NewAuthKey(os.Getenv("REDMINE_SERVER"), os.Getenv("REDMINE_KEY")).Trackers()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	t.Logf("%v\n", v)
}
