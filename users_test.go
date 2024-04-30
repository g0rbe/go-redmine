package redmine_test

import (
	"fmt"

	"gorbe.io/go/redmine"
)

func ExampleUsers() {

	users, err := redmine.NewPublic("https://www.redmine.org").Users()
	if err != nil {
		// handle error
	}

	fmt.Printf("%#v", users)
}

func ExampleUsers_withParams() {

	users, err := redmine.NewPublic("https://www.redmine.org").Users(redmine.NameParameter("tester"))
	if err != nil {
		// handle error
	}

	fmt.Printf("%#v", users)
}
