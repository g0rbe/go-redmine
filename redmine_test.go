package redmine_test

import (
	"gorbe.io/go/redmine"
)

func ExampleNewPublic() {

	_ = redmine.NewPublic("https://www.redmine.org")

}

func ExampleNewRegularLogin() {

	// admin1 user impersonates admin2
	_ = redmine.NewRegularLogin("https://www.redmine.org", "admin1", "password", "admin2")

}

func ExampleNewAuthKey() {

	// User of apikey impersonates admin2 if user is admin
	_ = redmine.NewAuthKey("https://www.redmine.org", "apikey", "admin2")

}

func ExampleNewHeaderKey() {

	// User of apikey impersonates admin2 if user is admin
	_ = redmine.NewHeaderKey("https://www.redmine.org", "apikey", "admin2")

}
